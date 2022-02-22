from config.config import ConnectionConfig
import time
import post
import generator
import logger
import math

config = ConnectionConfig()
config.HOST = 'http://127.0.0.1'
config.__dict__

log = logger.Log(__name__)

def retry_back_off(attempt):
    wait = math.pow(2, attempt)
    next = math.pow(2, attempt+1)
    time.sleep(wait)
    log.logger.warning("Next connection attempt will be performed in {} seconds".format(next))
    
def main():
    while True:
        array = generator.generate_array_of_random_integers()
        pay_load = {'elements': array}

        counter = 0
        
        err = post.health_request(**config.__dict__)
        while not err:
            retry_back_off(counter)
            err = post.health_request(**config.__dict__)
            counter += 1
        try:
            result = post.post_request(pay_load, **config.__dict__)
            log.logger.info('Shipped array: {}'.format(pay_load['elements']))
            time.sleep(1)
        except:
            log.logger.error('Remote end closed connection without response')


main()