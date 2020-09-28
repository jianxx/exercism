def classify(number):
    if number <= 0:
        raise ValueError("invalid number: %s", number)
    aliquot_sum = 0
    for x in range(1, number):
        if number % x == 0:
            aliquot_sum += x
    if aliquot_sum == number:
        return "perfect"
    elif aliquot_sum > number:
        return "abundant"
    else:
        return "deficient"
