def equilateral(sides):
    if any(side <= 0 for side in sides) or len(sides) != 3:
        return False
    return len(set(sides)) == 1


def isosceles(sides):
    if any(side <= 0 for side in sides) or len(sides) != 3:
        return False
    return (
        len(set(sides)) <= 2
        and sides[0] + sides[1] > sides[2]
        and sides[0] + sides[2] > sides[1]
        and sides[1] + sides[2] > sides[0]
    )


def scalene(sides):
    if any(side <= 0 for side in sides):
        return False
    return (
        len(set(sides)) == 3
        and sides[0] + sides[1] > sides[2]
        and sides[0] + sides[2] > sides[1]
        and sides[1] + sides[2] > sides[0]
    )
