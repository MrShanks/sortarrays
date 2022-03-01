import requests
import logger

log = logger.Log(__name__)


def query_string_format(**kwargs):
    return '{host}:{port}{endpoint}'.format(host = kwargs['host'], port = kwargs['port'], endpoint = kwargs['url'])


def health_request(**kwargs):
    log.logger.debug('Health request on address:{}'
                    .format(query_string_format(**dict(kwargs, **{'endpoint' : kwargs['health']}))))
    try:
        r = requests.get(query_string_format(**dict(kwargs, **{'endpoint' : kwargs['health']})))
    except requests.exceptions.RequestException as e:
        return
    return r


def post_request(pload, **kwargs):
    log.logger.debug('Posting on address: {}'
                    .format(query_string_format(**kwargs)))
    return requests.post(query_string_format(**kwargs), json=pload)