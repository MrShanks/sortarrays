import logging
import os


class Log:
    def __init__(self, module_name):
        # Create a custom logger
        self.logger = logging.getLogger(module_name)
        loglevel = os.environ.get('LOGLEVEL', 'INFO').upper()
        self.logger.setLevel(loglevel)

        # Create a custom handler
        self.stream_handler = logging.StreamHandler()

        # Create formatters and add it to handlers
        self.stream_formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
        self.stream_handler.setFormatter(self.stream_formatter)

        self.logger.addHandler(self.stream_handler)