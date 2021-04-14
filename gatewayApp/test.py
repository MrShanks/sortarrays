class Point(object):
    def __init__(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z

    @property
    def p(self):
        return [self.x, self.y, self.z]

    def __repr__(self):
        return 'x:{} y:{} z:{}'.format(self.x, self.y, self.z)

punto = Point(5,3,2)

print(punto)
print(punto.p)
print(punto.__repr__())