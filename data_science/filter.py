import numpy as np
from datetime import datetime


class Filtering:

    def __init__(self, data, answer):

        self.df = data.loc[data['rating'] <= 2000].copy()
        self.answer = answer

    def filter_interests(self, row, interests):

        for interest in interests:
            if row[interest] == 1:
                return 1
        return 0

    def get_time_flag(self, x):

        if x == '{"Пн":"","Вт":"","Ср":"","Чт":"","Пт":"","Сб":"","Вс":""}':
            return 1


        time_stamp = datetime.today()

        time_stamp = datetime(2022, 11, 5, 15, 17, 8, 132263)
        weekday = time_stamp.weekday()
        work_schedule = list(eval(x).values())[weekday]
        if len(work_schedule) == 0 or work_schedule is np.nan:
            return 0

        open_hour_minute = work_schedule[:5]
        close_hour_minute = work_schedule[-5:]

        time_closed = datetime(time_stamp.year, time_stamp.month, time_stamp.day,
                               int(close_hour_minute[:2]), int(close_hour_minute[-2:]))

        time_open = datetime(time_stamp.year, time_stamp.month, time_stamp.day,
                             int(open_hour_minute[:2]), int(open_hour_minute[-2:]))

        if time_open < time_stamp < time_closed and (time_closed - time_stamp).seconds >= 3600:
            return 1
        else:
            return 0

    def get_dataframe(self):

        if self.answer['animals'] == 1:
            self.df = self.df.loc[self.df['animals_flag'] == 1]
        if self.answer['kids'] == 1:
            self.df = self.df.loc[self.df['children_flag'] == 1]

        if self.answer['transport'] == '' or self.answer['transport'] is None or len(self.answer['transport']) == 0:
            pass

        if self.answer['transport'] == 'Общественный транспорт':
            self.df = self.df.loc[self.df['electrobus_flag'] == 1]

        elif self.answer['transport'] == 'Самокат':
            self.df = self.df.loc[self.df['scooter_flag'] == 1]

        elif self.answer['transport'] == 'Велосипед':
            self.df = self.df.loc[self.df['bike_flag'] == 1]

        if self.answer['with'] == '' or self.answer['with'] is None or len(self.answer['with']) == 0:
            pass

        elif self.answer['with'] == 'Пара':
            self.df = self.df.loc[self.df['date_flag'] == 1]

        elif self.answer['with'] == 'Семьей':
            self.df = self.df.loc[self.df['family_flag'] == 1]

        if self.answer['interests'] == '' or self.answer['interests'] is None or len(self.answer['interests']) == 0:
            self.df['ord'] = 200
        else:
            self.df.loc[self.df.apply(self.filter_interests, args=[self.answer['interests']], axis=1) == 1, 'ord'] = 100
            self.df['ord'].fillna(200)

        self.df['time_flag'] = self.df['work_schedule'].apply(self.get_time_flag)

        self.df.drop(self.df.loc[self.df['time_flag'] == 0].index, inplace=True)

        return self.df
