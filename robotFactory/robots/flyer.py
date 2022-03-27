from . import robot


class Flyer(robot.Robot):
    def __init__(self, name, id, *args):
        super().__init__(name, id, *args)
        self.type = 'flyer'

    def main_ability(self):
        print('Flyging')