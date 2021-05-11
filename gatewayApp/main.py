import time
import post
import generator
import logger


log = logger.Log(__name__)

HOST = 'http://sortarray'
PORT = '8080'
URL = '/api/v1/array/default'


def main():

    while True:
        array = generator.generate_array_of_random_integers()
        pay_load = {'elements': array}
        result = post.post_request('{}:{}{}'.format(HOST, PORT, URL), pay_load)
        log.logger.debug(f'posted {array}')
        log.logger.info(f'response from server: {result}')
        time.sleep(1)


main()
