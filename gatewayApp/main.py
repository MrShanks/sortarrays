import generateArray
import post


def main(path='http://127.0.0.1:5000/postArray', multiplier=1, array_size=3):
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

        res = post.post_request(path, pay_load)

    return res


main()