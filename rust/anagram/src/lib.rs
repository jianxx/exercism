use std::collections::HashMap;
use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &'a [&str]) -> HashSet<&'a str> {
    let word = word.to_lowercase();
    let word_map = word.chars().fold(HashMap::new(), |mut map, c| {
        *map.entry(c).or_insert(0) += 1;
        map
    });

    possible_anagrams
        .iter()
        .filter(|anagram| {
            if anagram.len() != word.len() {
                return false;
            }
            let anagram = anagram.to_lowercase();
            let anagram_map = anagram.chars().fold(HashMap::new(), |mut map, c| {
                *map.entry(c).or_insert(0) += 1;
                map
            });
            anagram != word && anagram_map == word_map
        })
        .map(|x| *x)
        .collect::<HashSet<&str>>()
}
