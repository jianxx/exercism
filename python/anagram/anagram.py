import itertools
from string import ascii_lowercase


def find_anagrams(word, candidates):
    charcount = [0] * 26

    for c in word.lower():
        if c in ascii_lowercase:
            charcount[ord(c) - ord("a")] += 1

    def is_anagram(candidate):
        if len(word) != len(candidate):
            return False
        candidate_count = [0] * 26
        for c in candidate.lower():
            if c in ascii_lowercase:
                candidate_count[ord(c) - ord("a")] += 1
        return word.lower() != candidate.lower() and all(
            charcount[i] == candidate_count[i] for i in range(len(charcount))
        )

    return itertools.filterfalse(
        lambda candidate: not is_anagram(candidate), candidates
    )

