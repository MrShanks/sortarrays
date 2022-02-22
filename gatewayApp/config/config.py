class ConnectionConfig:
    def __init__(self,
                 host = 'http://sortarray',
                 port = '8080',
                 url = '/api/v1/array/default',
                 health = '/health'
        ):
        self.HOST = host
        self.PORT = port
        self.URL = url
        self.HEALTH = health

    def get_attr(self):   
        return [i for i in self.__dict__.keys() if i[:1] != '_']