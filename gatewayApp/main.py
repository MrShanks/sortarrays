import generateArray
import post


def main(multiplier=1, max_array_size=3, path='http://127.0.0.1:5000/postArray'):
    while True:
        array = generateArray.generate_array(multiplier, max_array_size)
        pay_load = {'vector': array}
        res = post.post_request(path, pay_load)

    return res