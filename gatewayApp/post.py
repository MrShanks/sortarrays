import requests
import logger

log = logger.Log(__name__)


def post_request(path, pload):
    log.logger.debug('post is posting')
    return requests.post(path, json=pload)