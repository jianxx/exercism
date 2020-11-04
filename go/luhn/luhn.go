package luhn

// Valid determines whether or not a number is valid per the Luhn formula.
func Valid(number string) bool {
	digitCount := 0
	checkSum := 0
	for i := len(number) - 1; i >= 0; i-- {
		switch {
		case number[i] == ' ':
			continue
		case number[i] >= '0' && number[i] <= '9':
			n := int(number[i] - '0') // iterator value
			if digitCount%2 != 0 {
				n <<= 1
				if n > 9 {
					n -= 9
				}
			}
			checkSum += n
			digitCount++
		default:
			return false
		}
	}
	return digitCount > 1 && checkSum%10 == 0
}
