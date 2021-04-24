import logging


class Log:
    def __init__(self, module_name):
        # Create a custom logger
        self.logger = logging.getLogger(module_name)
        self.logger.setLevel(logging.INFO)

        # Create a custom handler
        self.stream_handler = logging.StreamHandler()

        # Create formatters and add it to handlers
        self.stream_formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
        self.stream_handler.setFormatter(self.stream_formatter)

        self.logger.addHandler(self.stream_handler)