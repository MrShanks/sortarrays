import http
import time
import post
import generator
import logger
import math

log = logger.Log(__name__)

config = dict(
         HOST = 'http://127.0.0.1',
         PORT = '8080',
         ENDPOINT = '',
         URL = '/api/v1/array/default',
         HEALTH = '/health'
)

def retry_back_off(attempt):
    wait = math.pow(2, attempt)
    next = math.pow(2, attempt+1)
    time.sleep(wait)
    log.logger.warning("Next connection attempt will be performed in {} seconds".format(next))

def query_string_format(HOST = config['HOST'], PORT = config['PORT'], ENDPOINT = config['ENDPOINT']):
    return '{host}:{port}{endpoint}'.format(host = HOST, port = PORT, endpoint = ENDPOINT)
    
def main():
     while True:
        array = generator.generate_array_of_random_integers()
        pay_load = {'elements': array}

        print()
        counter = 0
        
        err = post.request_health(query_string_format(ENDPOINT = config['HEALTH']))
        while not err:
            retry_back_off(counter)
            err = post.request_health(query_string_format(ENDPOINT = config['HEALTH']))
            counter += 1
        try:
            result = post.post_request(query_string_format(ENDPOINT = config['URL']), pay_load)
            log.logger.info('Shipped array: {}'.format(pay_load['elements']))
            time.sleep(1)
        except:
            log.logger.error('Remote end closed connection without response')


main()