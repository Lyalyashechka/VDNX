import psycopg2

from place import Place


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
