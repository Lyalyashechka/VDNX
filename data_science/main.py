import json

import psycopg2

from flask import Flask, request

from db import DB

db = DB()
app = Flask(__name__)


@app.route("/routes/personal", methods=["POST"])
def personal_routes():
    places = db.get_all_places()
    personal_info = request.get_json(force=True)
    print(personal_info)

    return json.dumps({"places": places})


if __name__ == "__main__":
    app.run()
