from flask import Flask, Response, jsonify, render_template, request
from flask_cors import CORS, cross_origin
import time

app = Flask(__name__)
cors = CORS(app)
app.config['CORS_HEADERS'] = 'Content-Type'

count = 0


def event_stream():
    def generate():
        while True:
            time.sleep(1)
            yield f"data: {count}\n\n"

    return Response(generate(),
                    200,
                    mimetype="text/event-stream",
                    )


@app.get("/")
def home():
    return render_template("index.html")


@app.get("/events")
@cross_origin()
def events():
    return event_stream()


# DRIVER RELATED ENDPOINTS
@app.post("/sos")
def sos():
    data = request.json
    global count
    count += 1
    return jsonify(data)


if __name__ == '__main__':
    app.run(debug=True)
