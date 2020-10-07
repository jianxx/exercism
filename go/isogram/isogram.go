package isogram

import "unicode"

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(word string) bool {
	charSet := make(map[rune]bool)
	for _, c := range word {
		c := unicode.ToLower(c)
		if c >= 'a' && c <= 'z' {
			exists := charSet[c]
			if exists {
				return false
			}
			charSet[c] = true
		}
	}
	return true
}
