from flask import Flask

#from .utils import generateArray


app = Flask(__name__)

@app.route('/postArray', methods = ['POST'])
def postArray():
    return 'Postato'


@app.route('/getArray')
def getArray():
    return str(generateArray.generateArray())