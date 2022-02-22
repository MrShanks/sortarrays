from config.config import ConnectionConfig
import time
import post
import generator
import logger
import math

CONFIG = ConnectionConfig('http://sortarray', '8080', '/api/v1/array/default', '/health')

BASE_URL = '{}:{}'.format(CONFIG.get_host(), CONFIG.get_port())
ARRAY_URL = '{}{}'.format(BASE_URL, CONFIG.get_endpoint())
HEALTH_URL = '{}{}'.format(BASE_URL, CONFIG.get_health_endpoint())

log = logger.Log(__name__)


def retry_back_off(attempt):
    wait = math.pow(2, attempt)
    next_attempt = math.pow(2, attempt + 1)
    time.sleep(wait)
    log.logger.warning("Next connection attempt will be performed in {} seconds".format(next_attempt))


def main():
    while True:
        array = generator.generate_array_of_random_integers()
        pay_load = {'elements': array}

        counter = 0

        err = post.health_request(HEALTH_URL)
        while not err:
            retry_back_off(counter)
            err = post.health_request(HEALTH_URL)
            counter += 1

        try:
            post.post_request(pay_load, ARRAY_URL)
            log.logger.info('Shipped array: {}'.format(pay_load['elements']))
            time.sleep(1)
        except Exception as e:
            log.logger.error('Remote end closed connection without response with error: {}'.format(e))


if __name__ == "__main__":
    main()
