from . import robot


class Singer(robot.Robot):
    def __init__(self, name, id, *args):
        super().__init__(name, id, *args)
        self.type = 'singer'

    def main_ability(self):
        print('Singing')