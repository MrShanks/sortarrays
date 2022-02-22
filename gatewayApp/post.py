import requests
import logger

log = logger.Log(__name__)

def query_string_format(**kwargs):
    return '{host}:{port}{endpoint}'.format(host = kwargs['HOST'], port = kwargs['PORT'], endpoint = kwargs['URL'])

def health_request(**kwargs):
    log.logger.debug('Health request on address:{}'
                    .format(query_string_format(**dict(kwargs, **{'URL' : kwargs['HEALTH']}))))
    try:
        r = requests.get(query_string_format(**dict(kwargs, **{'URL' : kwargs['HEALTH']})))
    except requests.exceptions.RequestException as e:
        return 
    return r

def post_request(pload, **kwargs):
    log.logger.debug('Posting on address: {}'
                    .format(query_string_format(**kwargs)))
    response = requests.post(query_string_format(**kwargs), json=pload)
    return response