import datetime
import time
import post
import generator
import logger
import math
import yaml
import os
import multiprocessing
import check_status_gateway

log = logger.Log(__name__)

with open(os.path.dirname(os.path.abspath(__file__)) + '/config/config.yaml') as file:
    config = yaml.load(file, Loader=yaml.FullLoader)


def retry_back_off(attempt):
    wait = math.pow(2, attempt)
    next_attempt = math.pow(2, attempt + 1)
    time.sleep(wait)
    log.logger.warning("Next connection attempt will be performed in {} seconds".format(next_attempt))


def difference_between_two_times(t1,
                                 t2 = None,
                                 format1 = '%a, %d %b %Y %H:%M:%S %Z',
                                 secconds = True):
    t2 = t2 if t2 != None else datetime.datetime.utcnow()
    deltat = datetime.datetime.strptime(t1, format1) - t2
    return deltat.total_seconds() if secconds else deltat


def main():
    check_status()


def check_status():
    counter = 0
    status_code = post.health_request(config)
    while status_code != 200:
        retry_back_off(counter)
        status_code = post.health_request(config)
        counter += 1
    send()


def send():
    post.signup_request(config)
    auth = post.login_request(config)
    missing_time = difference_between_two_times(auth['Set-Cookie'].split('Expires=')[1])

    pay_load = {'elements': generator.generate_array_of_random_integers()}

    status_code = post.post_request(pay_load, auth['Set-Cookie'].split(';')[0], config)
    while status_code == 201:
        if missing_time < 25:
            auth = post.refresh_request(auth['Set-Cookie'].split(';')[0], config)
        missing_time = difference_between_two_times(auth['Set-Cookie'].split('Expires=')[1])
        pay_load = {'elements': generator.generate_array_of_random_integers()}
        try:
            status_code = post.post_request(pay_load, auth['Set-Cookie'].split(';')[0], config)
            log.logger.info('Shipped array: {}'.format(pay_load['elements']))
            time.sleep(1)
        except Exception as e:
            log.logger.error('Remote end closed connection without response with error {}'.format(e))
            check_status()


if __name__ == "__main__":
    proces1 = multiprocessing.Process(target = main)
    proces2 = multiprocessing.Process(target = check_status_gateway.start_server)
    proces1.start()
    proces2.start()