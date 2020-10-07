import re


def is_isogram(string):
    word = re.sub("[^a-z]", "", string.lower())
    return len(word) == len(set(word))
