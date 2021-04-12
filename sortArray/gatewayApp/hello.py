from flask import Flask

from utils import generateArray


app = Flask(__name__)

@app.route('/postArray')
def postArray():
    return generateArray.generateArray()
