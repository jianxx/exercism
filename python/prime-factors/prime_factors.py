import math


def factors(value):
    r = []
    while value % 2 == 0:
        r.append(2)
        value /= 2
    for x in range(3, int(math.sqrt(value)) + 1, 2):
        if value % x == 0:
            r.append(x)
            value /= x
    if value > 2:
        r.append(value)
    return r

