class ConnectionConfig:
    def __init__(self, host, port, endpoint, health_endpoint):
        self.host = host
        self.port = port
        self.endpoint = endpoint
        self.health_endpoint = health_endpoint

    def get_host(self):
        return self.host

    def get_port(self):
        return self.port

    def get_endpoint(self):
        return self.endpoint

    def get_health_endpoint(self):
        return self.health_endpoint

    def set_host(self, host):
        self.host = host

    def set_port(self, port):
        self.port = port

    def set_endpoint(self, endpoint):
        self.endpoint = endpoint

    def set_health_endpoint(self, health_endpoint):
        self.health_endpoint = health_endpoint
