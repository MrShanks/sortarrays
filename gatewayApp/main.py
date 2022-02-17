import http
import time
import post
import generator
import logger
import math

log = logger.Log(__name__)

config = dict(
         HOST = 'http://sortarray',
         PORT = '8080',
         URL = '/api/v1/array/default',
         HEALTH = '/health'
)

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
        
        err = post.request_health('{}:{}{}'.format(config['HOST'], config['PORT'], config['HEALTH']))
        while not err:
            retry_back_off(counter)
            err = post.request_health('{}:{}{}'.format(config['HOST'], config['PORT'], config['HEALTH']))
            counter += 1
        try:
            result = post.post_request('{}:{}{}'.format(config['HOST'], config['PORT'], config['URL']), pay_load)
            log.logger.info('Shipped array: {}'.format(pay_load['elements']))
            time.sleep(1)
        except:
            log.logger.error('Remote end closed connection without response')


main()