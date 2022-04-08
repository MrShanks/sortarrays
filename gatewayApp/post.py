import requests
import logger

log = logger.Log(__name__)


def query_string_format(config, endpoint):
    return '{host}:{port}{endpoint}'.format(host = config['host'], port = config['port'], endpoint = endpoint)


def health_request(config):
    health_url = query_string_format(config, '/health')
    log.logger.debug('Health request on address:{}'.format(health_url))
    try:
        r = requests.get(health_url)
    except requests.exceptions.RequestException as e:
        return
    return r.status_code


def post_request(pload, token, config):
    post_url = query_string_format(config, '/api/v1/array/default')
    log.logger.debug('Posting on address: {}'.format(post_url))
    return requests.post(post_url, headers={'Cookie' : token, 'Content-Type': 'text/plain'}, json=pload).status_code


def login_request(pload, config):
    login_url = query_string_format(config, '/signin')
    log.logger.debug('Getting token: {}'.format(login_url))
    return requests.post(login_url, json=pload).headers