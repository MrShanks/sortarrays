import random as rd


def generate_array(multiplier=1, max_array_size=3):
    return [rd.random() * multiplier for x in range(max_array_size)]