from flask import Flask

app = Flask(__name__)


@app.route("/")
def hello():
    return "<p style='color:red'>Welcome on the gatewayapp main page!</p>"


@app.route("/health")
def health_status():
    return "status: ok"


def start_server():
    app.run(host='127.0.0.1', port=3000, debug=True)
    return None