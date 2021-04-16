import json

import generateArray
import post


def main(path='http://localhost:8080/api/v1/array/default', multiplier=1, array_size=3):
    """
    generate an array and send it to the passed server

    Args:
        path: path where to make post request
        multiplier: this is multiplier for array elements, by default they are between 0-1
        array_size: elements of the array

    Returns: response from server

    """
    while True:
        array = generateArray.generate_array(multiplier, array_size)
        pay_load = {'vector': array}
        print(pay_load)
        res = post.post_request(path, pay_load)
        print(res)

    return res


main()