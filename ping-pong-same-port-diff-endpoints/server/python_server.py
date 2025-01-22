from flask import Flask

app = Flask(__name__)

@app.route('/python-ping', methods=['GET'])
def python_ping():
    return "pong from python", 200

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8082)
