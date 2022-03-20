class Robot:
    def __init__(self, name, id, *args):
        self.name = name
        self.id = id
        self.type = 'base'
        self.drive_ability = 'drive'


    def drive(self):
        print('Driving')


    def main_ability(self):
        print('')


    def get_properties(self):
        return self.name, self.id, self.type