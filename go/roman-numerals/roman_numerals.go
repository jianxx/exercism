package romannumerals

import (
	"bytes"
	"errors"
	"sort"
)

// ToRomanNumeral converts normal numbers to Roman Numerals
func ToRomanNumeral(arabic int) (string, error) {
	if arabic <= 0 || arabic > 3000 {
		return "", errors.New("number is out of range")
	}
	var m = map[int]string{1000: "M", 900: "CM", 500: "D", 400: "CD", 100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}
	romanNumbers := make([]int, 0, len(m))
	for k := range m {
		romanNumbers = append(romanNumbers, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(romanNumbers)))
	var b bytes.Buffer
	for _, k := range romanNumbers {
		if arabic < k {
			continue
		}
		for arabic >= k {
			arabic -= k
			b.WriteString(m[k])
		}
	}
	return b.String(), nil
}
