from string import ascii_lowercase


def is_pangram(sentence):
    chars = set()
    for c in list(sentence.lower()):
        if c in list(ascii_lowercase):
            chars.add(c)
    return len(chars) == 26
