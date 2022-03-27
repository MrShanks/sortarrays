import robots

id = 0
class Factory:
    def __init__(self) -> None:
        pass


    def create_robot(self, name, type = 'base'):
        global id
        id += 1
        if type == 'sing':
            return robots.singer.Singer(name, id)
        elif type == 'fly':
            return robots.flyer.Flyer(name, id)
        elif type == 'clean':
            return robots.cleaner.Cleaner(name, id)
        else:
            return robots.robot.Robot(name, id)