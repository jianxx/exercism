def is_paired(input_string):
    l = []
    switcher = {
        "[": lambda c: l.append(c) or True,
        "(": lambda c: l.append(c) or True,
        "{": lambda c: l.append(c) or True,
        "]": lambda c: len(l) > 0 and l.pop() == "[",
        ")": lambda c: len(l) > 0 and l.pop() == "(",
        "}": lambda c: len(l) > 0 and l.pop() == "{",
    }
    for c in input_string:
        method = switcher.get(c, lambda c: True)
        if not method(c):
            return False

    return len(l) == 0
