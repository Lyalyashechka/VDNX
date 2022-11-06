import pandas.io.sql as psql
import psycopg2

from model import RouteModel, NpEncoder
from filter import Filtering
from flask import Flask, request, json, jsonify


class DB:
    def __init__(self):
        self.conn = psycopg2.connect(dbname='vdnx', user='postgres',
                                     password='1234', host='localhost')
        self.cursor = self.conn.cursor()

    def get_df(self):
        self.cursor.execute("SELECT * from places")
        self.conn.commit()

        df = psql.read_sql('SELECT * FROM places', self.conn)

        return df



db = DB()
app = Flask(__name__)

@app.route("/routes/personal", methods=['POST'])
def hello():
    answer = request.get_json(force=True)

    #answer = {'with': 'Компанией',
    #          'animals': 1,
    #          'kids': 1,
    #          'interests': ['promenade'],
    #          'transport': 'Общественный транспорт'
    #          }

    df = db.get_df()

    my_filter = Filtering(df, answer)
    selected_df = my_filter.get_dataframe()

    if selected_df.shape[0] < 1:
        return []
    model = RouteModel(selected_df, df, answer)
    routes = model.get_routes()
    return json.dumps({"routes": routes}, cls=NpEncoder)


if __name__ == "__main__":
    app.run()
