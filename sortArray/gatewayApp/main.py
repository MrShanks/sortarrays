import requests
from utils import generateArray

#r = requests.get('http://127.0.0.1:5000/postArray')

tmpArray = generateArray.generateArray()

pload = {'vector':tmpArray}

r = requests.post('http://127.0.0.1:5000/postArray', data = pload)


print(r.text)