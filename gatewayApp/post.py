import requests


def post_request(path, pload):
    """

    Args:
        path:
        pload:

    Returns:

    """
    result = requests.post(path, json=pload)
    return result