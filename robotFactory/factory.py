import robot, singer, flyer, cleaner

id = 0

class Factory:
    def __init__(self) -> None:
        pass


    def create_robot(self, name, type = 'base'):
        global id
        id += 1
        if type == 'sing':
            return singer.Singer(name, id)
        elif type == 'fly':
            return flyer.Flyer(name, id)
        elif type == 'clean':
            return cleaner.Cleaner(name, id)
        else:
            return robot.Robot(name, id)