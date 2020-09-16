package scrabble

import (
	"strings"
)

// Score : given a word, compute the Scrabble score for that word.
func Score(input string) int {
	var values = [26]int{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}
	var score = 0
	for _, c := range strings.ToLower(input) {
		score += values[c-'a']
	}
	return score
}
