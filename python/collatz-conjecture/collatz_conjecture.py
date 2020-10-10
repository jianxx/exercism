def steps(number):
    if number <= 0:
        raise ValueError("Invalid input: %d", number)
    steps = 0
    while number > 1:
        number = number >> 1 if number % 2 == 0 else number * 3 + 1
        steps += 1
    return steps
