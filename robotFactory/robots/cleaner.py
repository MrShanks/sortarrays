from . import robot


class Cleaner(robot.Robot):
    def __init__(self, name, id, *args):
        super().__init__(name, id, *args)
        self.type = 'cleaner'

    def main_ability(self):
        print('Cleaning')