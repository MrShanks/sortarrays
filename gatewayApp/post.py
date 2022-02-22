import requests
import logger

log = logger.Log(__name__)


def health_request(config):
    log.logger.debug('Health request on address:{}'.format(config))
    try:
        r = requests.get(config)
    except requests.exceptions.RequestException as e:
        return e
    return r


def post_request(payload, config):
    log.logger.debug('Posting on address: {}'.format(config))
    return requests.post(config, json=payload)
