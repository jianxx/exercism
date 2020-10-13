package reverse

func Reverse(input string) string {
	n := len(input)
	runes := make([]rune, n)
	for _, rune := range input {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}
