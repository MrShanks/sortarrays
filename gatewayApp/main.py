import time
import post
import generate_array_of_random_integers
import logger


log = logger.Log(__name__)


def main():
    endpoint = 'http://localhost:8080/api/v1/array/default'

    while True:
        array = generate_array_of_random_integers.generate_array()
        pay_load = {'elements': array}
        result = post.post_request(endpoint, pay_load)
        log.logger.debug(f'posted {array}')
        log.logger.info(f'response from server: {result}')
        time.sleep(1)


main()