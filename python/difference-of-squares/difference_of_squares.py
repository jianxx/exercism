def square_of_sum(number):
    return sum(range(number + 1)) ** 2


def sum_of_squares(number):
    return sum(map(lambda n: n ** 2, range(number + 1)))


def difference_of_squares(number):
    return square_of_sum(number) - sum_of_squares(number)
