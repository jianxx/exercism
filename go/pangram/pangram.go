package pangram

import "unicode"

func IsPangram(sentence string) bool {
	charSet := make(map[rune]bool)
	for _, c := range sentence {
		c := unicode.ToLower(c)
		if c >= 'a' && c <= 'z' {
			charSet[c] = true
		}
	}

	return len(charSet) == 26
}
