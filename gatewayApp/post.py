import requests

def post_request(path='http://127.0.0.1:5000/postArray', pload):
    return requests.post(path, data=pload)