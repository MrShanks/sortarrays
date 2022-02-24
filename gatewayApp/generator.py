import random as rd
import logger

log = logger.Log(__name__)


def generate_array_of_random_integers(min_value=0, max_value=9, array_size=3):
    log.logger.debug('Generating new array')
    return [rd.randint(min_value, max_value) for x in range(array_size)]