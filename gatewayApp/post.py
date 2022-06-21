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


def signup_request(config):
    signup_url  = query_string_format(config, '/signup')
    log.logger.debug('Signup to Sortarray')
    return requests.post(signup_url , json={key: config.get(key) for key in ['username', 'password']}).headers


def login_request(config):
    login_url = query_string_format(config, '/signin')
    log.logger.debug('Getting token from: {}'.format(login_url))
    return requests.post(login_url, json={key: config.get(key) for key in ['username', 'password']}).headers


def refresh_request(token, config):
    refresh_url = query_string_format(config, '/refresh')
    log.logger.debug('Getting new token from: {}'.format(refresh_url))
    return requests.post(refresh_url, headers={'Cookie' : token, 'Content-Type': 'text/plain'}).headers