import json
import numpy as np
import pandas as pd
import warnings
from scipy.spatial import distance

warnings.filterwarnings("ignore")


class NpEncoder(json.JSONEncoder):

    def default(self, obj):
        if isinstance(obj, np.integer):
            return int(obj)
        if isinstance(obj, np.floating):
            return float(obj)
        if isinstance(obj, np.ndarray):
            return obj.tolist()
        return super(NpEncoder, self).default(obj)


class RouteModel:

    def __init__(self, filtered, df, answer):

        self.filtered = filtered
        self.df = df
        self.answer = answer

    def get_closest_services(self, route, service_name, n_services=10):

        result = self.df.loc[self.df['type_service'] == service_name].copy()
        min_distance = []
        for service in result['coordinates'].values:
            distances = [distance.cosine(service, point) for point in
                         route.loc[route['rating'] <= 2000]['coordinates'].values]
            min_distance.append(min(distances))
        result['distance'] = min_distance

        return result.sort_values(by='distance', ascending=True).head(n_services)

    def get_closest_point(self, route, recommend):

        recommend = recommend.loc[(~recommend['coordinates'].isin(route['coordinates'].values)) &
                                  (~recommend['title'].isin(route['title'].values))]

        vectors_to_compare = ['vector_title', 'vector_text', 'coordinates']

        for vector_to_compare in vectors_to_compare:
            mean_vector = np.mean(
                np.array([np.array(vector) for vector in route[vector_to_compare]]), axis=0)

            recommend['distance_' + vector_to_compare] = [distance.cosine(mean_vector, point_vector)
                                                          for point_vector in recommend[vector_to_compare].values]
        distance_coef = 10 ** 2
        title_coef = 0.1
        text_coef = 1

        if route.loc[route['type_place'].isin(['Детская площадка', 'Зеленая зона',
                                               'Фонтаны'])].shape[0] / route.shape[0] >= 0.5:
            distance_coef = 10 ** 10
            title_coef = 0
            text_coef = 0.001

        if self.answer['transport'] in ['Велосипед', 'Общественный транспорт']:
            distance_coef = 0.2

        elif self.answer['transport'] == 'Самокат':
            distance_coef = 0.5

        recommend['distance'] = recommend['distance_vector_title'] * title_coef \
                                + recommend['distance_vector_text'] * text_coef \
                                + recommend['distance_coordinates'] * distance_coef

        if route.loc[route['type_place']=='Детская площадка'].shape[0] / route.shape[0] >= 0.4:
            recommend.loc[recommend['type_place'] == 'Детская площадка', 'ord'] *= 5

        recommend.sort_values(by=['ord', 'distance'], ascending=True, inplace=True)

        return recommend.head(1)

    def get_route(self, corpus, n_points=10):

        route = corpus.sort_values(by=['ord', 'rating']).head(5).sample(1)

        for _ in range(n_points - 1):
            route = pd.concat([route, self.get_closest_point(route, corpus)])

        if self.answer['transport'] == 'Самокат':
            route = pd.concat(
                [route, self.get_closest_services(route, 'Прокат', n_services=8)])
        elif self.answer['transport'] == 'Общественный транспорт':
            route = pd.concat([route, self.get_closest_services(
                route, 'Остановка', n_services=8)])

        route = pd.concat([route, self.get_closest_services(
            route, service_name='Еда', n_services=10)])
        route = pd.concat([route, self.get_closest_services(
            route, service_name='Туалеты', n_services=10)])

        return route

    def get_routes_raw(self):

        raw_routes = []
        tmp = self.filtered.copy()

        if self.answer['transport'] == 'Велосипед':
            n_points = [10]
            proba = [1.0]
        elif self.answer['animals'] == 1:
            n_points = [8]
            proba = [1.0]
        else:
            n_points = [3, 4]
            proba = [0.2, 0.8]

        while len(raw_routes) < 3:

            if tmp.shape[0] < max(n_points) // 2:
                break
            if tmp.shape[0] <= 6:
                n_points = [3]
                proba = [1.0]

            full_route = self.get_route(tmp, n_points=np.random.choice(n_points, p=proba))

            raw_routes.append(full_route)
            tmp = tmp[~tmp['id'].isin(
                full_route.loc[full_route['rating'] <= 2000]['id'].values)]

        return raw_routes

    def get_main_points_sorted(self, df):

        df['my_distance'] = df['coordinates'].apply(
            lambda x: distance.cosine(x, [55.8262103, 37.63772804]))
        df.sort_values(by='my_distance', inplace=True)

        return df['json'].values

    def get_routes(self):

        route_list = self.get_routes_raw()
        final = []
        for i, route in enumerate(route_list):
            if self.answer['transport'] == 'Велосипед' or self.answer['animals'] == 1:
                duration = route.loc[(route['rating'] <= 2000) &
                                     (~route['type_place'].isin(['Детская площадка', 'Зеленая зона']))].shape[0] * 5 + \
                           route.loc[route['type_place'].isin(['Детская площадка', 'Зеленая зона'])]['duration'].sum()
            else:
                duration = route['duration'].sum()

            route.drop(columns=['vector_text', 'vector_title',
                                'time_flag', 'distance_vector_title', 'distance_vector_text',
                                'distance_coordinates', 'distance'], inplace=True)
            route['json'] = route.apply(lambda x: json.loads(x.to_json()), axis=1)

            final.append({
                'id': str(i),
                'main': self.get_main_points_sorted(route.loc[route['rating'] <= 2000]),
                'sup_points': route.loc[route['rating'] > 2000]['json'].values,
                'duration': duration
            }
            )
        return final
