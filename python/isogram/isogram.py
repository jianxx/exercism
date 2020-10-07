from string import ascii_lowercase


def is_isogram(string):
    charset = set()
    for c in string.lower():
        if c in ascii_lowercase:
            if c in charset:
                return False
            charset.add(c)
    return True
