import requests


def post_request(path, pload):
    """

    Args:
        path:
        pload:

    Returns:

    """
    result = requests.post(path, pload)
    return result