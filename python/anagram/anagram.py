from collections import Counter


def find_anagrams(word, candidates):
    word_counter = Counter(word.lower())

    return [
        candidate
        for candidate in candidates
        if word.lower() != candidate.lower()
        and Counter(candidate.lower()) == word_counter
    ]
