from flask import Flask

#from .utils import generate_array


app = Flask(__name__)

@app.route('/postArray', methods = ['POST'])
def postArray():
    return 'Postato'


@app.route('/getArray')
def getArray():
    return str(generateArray.generate_array())