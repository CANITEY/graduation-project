from flask import Flask, Response, jsonify, render_template, request
import time

app = Flask(__name__)

count = 0


def event_stream():
    def generate():
        while True:
            time.sleep(.5)
            yield f"data: {count}\n\n"

    return Response(generate(), 200, mimetype="text/event-stream")


@app.get("/")
def home():
    return render_template("index.html")


@app.get("/ss")
def ss():
    return "hehe"


@app.get("/events")
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
