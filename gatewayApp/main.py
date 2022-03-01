import time
import post
import generator
import logger
import math
import yaml
import os


def retry_back_off(attempt):
    wait = math.pow(2, attempt)
    next_attempt = math.pow(2, attempt + 1)
    time.sleep(wait)
    log.logger.warning("Next connection attempt will be performed in {} seconds".format(next_attempt))


def main():
    log = logger.Log(__name__)

    with open(os.path.dirname(os.path.abspath(__file__)) + '/config/config.yaml') as file:
        config = yaml.load(file, Loader=yaml.FullLoader)
        
    while True:
        pay_load = {'elements': generator.generate_array_of_random_integers()}

        counter = 0
        
        err = post.health_request(**config)
        while not err:
            retry_back_off(counter)
            err = post.health_request(**config)
            counter += 1
        try:
            post.post_request(pay_load, **config)
            log.logger.info('Shipped array: {}'.format(pay_load['elements']))
            time.sleep(1)
        except Exception as e:
            log.logger.error('Remote end closed connection without response with error {}'.format(e))


if __name__ == "__main__":
    main()