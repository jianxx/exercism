def score(x, y):
    if x * x + y * y <= 1:
        return 10
    elif x * x + y * y <= 25 and x * x + y * y > 1:
        return 5
    elif x * x + y * y <= 100 and x * x + y * y > 25:
        return 1
    else:
        return 0

