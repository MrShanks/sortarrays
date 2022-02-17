import requests
import logger

log = logger.Log(__name__)

def request_health(path):
    try:
        r = requests.get(path)
    except requests.exceptions.RequestException as e:
        return 
    return r

def post_request(path, pload):
    log.logger.debug('post is posting')
    response = requests.post(path, json=pload)
    return response