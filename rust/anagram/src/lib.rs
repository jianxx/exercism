use std::collections::HashMap;
use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &'a [&str]) -> HashSet<&'a str> {
    let word = word.to_lowercase();
    let anagrams = word.chars().fold(HashMap::new(), |mut anagrams, c| {
        *anagrams.entry(c).or_insert(0) += 1;
        anagrams
    });

    possible_anagrams
        .iter()
        .filter(|anagram| {
            if anagram.len() != word.len() {
                return false;
            }
            let anagram = anagram.to_lowercase();
            let charmap = anagram.chars().fold(HashMap::new(), |mut anagrams, c| {
                *anagrams.entry(c).or_insert(0) += 1;
                anagrams
            });
            anagram != word && charmap == anagrams
        })
        .map(|x| *x)
        .collect::<HashSet<&str>>()
}
