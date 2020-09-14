def is_armstrong_number(number):
    result = 0
    number_str = str(number)
    for c in number_str:
        result += pow(int(c), len(number_str))
    return result == number
