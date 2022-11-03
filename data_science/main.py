import json

import psycopg2

from flask import Flask


class Place:
    def __init__(self, id, coord, title):
        self.id = id
        self.coord = coord
        self.title = title

    def to_json(self):
        return json.dumps(self, default=lambda o: o.__dict__)


class DB:
    def __init__(self):
        self.conn = psycopg2.connect(dbname='vdnx', user='postgres',
                                password='1234', host='localhost')
        self.cursor = self.conn.cursor()

    def get_all_places(self):
        self.cursor.execute("SELECT * from places where is_event = false")
        self.conn.commit()
        rows = self.cursor.fetchall()

        places = []

        for row in rows:
            place = Place(row[0], row[3], row[14])
            places.append(place.to_json())

        return places


db = DB()
app = Flask(__name__)


@app.route("/routes/personal")
def hello():
    places = db.get_all_places()
    return json.dumps({"places": places})


if __name__ == "__main__":
    app.run()
