from flask import Flask

app = Flask(__name__)

@app.route('/postArray', methods = ['POST', 'GET'])
def postArray():
    return 'Postato'


@app.route('/getArray')
def getArray():
    return "GET WORKS"