import json


class Place:
    def __init__(self, id, coord, title):
        self.id = id
        self.coord = coord
        self.title = title

    def to_json(self):
        return json.dumps(self, default=lambda o: o.__dict__)

