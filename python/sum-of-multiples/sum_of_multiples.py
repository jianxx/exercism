def sum_of_multiples(limit, multiples):
    return sum({x for i in multiples if i > 0 for x in range(i, limit, i)})
