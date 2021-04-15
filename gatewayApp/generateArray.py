import random as rd


def generate_array(multiplier=1, array_size=3):
    """
    
    Args:
        multiplier: this is multiplier for array elements, by default they are between 0-1
        array_size: elements of the array

    Returns:

    """
    return [rd.random() * multiplier for x in range(array_size)]