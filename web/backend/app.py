from flask import Flask, jsonify, render_template, request
from flask_cors import CORS, cross_origin

from stream import Stream

app = Flask(__name__)
cors = CORS(app)
app.config['CORS_HEADERS'] = 'Content-Type'
stream = Stream(app)


@app.get("/")
def home():
    return render_template("index.html")


@app.get("/events")
@cross_origin()
def events():
    return stream.eventStream(stream.publish)


# DRIVER RELATED ENDPOINTS
@app.post("/sos")
def sos():
    data: dict = request.json
    carInfo = {}
    try:
        carInfo = {
                "carUUID": data["UUID"],
                "longitude": data["longitude"],
                "latitude": data["latitude"]
                }
    except KeyError:
        resp = {
                "status": "failure"
                }
        return jsonify(resp)

    resp = {
            "status": "sucess"
            }

    # TODO: add the carInfo as a json to buffer append and try to
    # send it as a json
    stream.buffer.append(jsonify(carInfo))
    events()
    return jsonify(resp)


if __name__ == '__main__':
    app.run(debug=True)
