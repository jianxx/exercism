import itertools
from string import ascii_lowercase


def is_isogram(string):
    charcount = [0] * 26

    def addchar(c):
        if c in ascii_lowercase:
            charcount[ord(c) - ord("a")] += 1

    itertools.starmap(lambda c: addchar(c), string.lower())
    return all(count < 2 for count in charcount)
