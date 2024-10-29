import time
from flask import Flask, Response


class Stream:
    def __init__(self, app: Flask):
        self.count = 0
        self.buffer = []
        self.app = app

    def eventStream(self, callback, message=None):
        return Response(callback(),
                        200,
                        mimetype="text/event-stream",
                        )

    def generate(self):
        yield f'data: {{"count": {self.count}}}\n\n'

    def publish(self):
        yield f'data: {self.buffer}\n\n'
        try:
            self.buffer = []
        except:
            pass
