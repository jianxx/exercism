package luhn

func Valid(number string) bool {
	digitCount := 0
	checkSum := 0
	for i := len(number) - 1; i >= 0; i-- {
		if number[i] == ' ' {
			continue
		} else if number[i]-'0' > 9 {
			return false
		}
		if digitCount%2 != 0 {
			if (number[i]-'0')*2 > 9 {
				checkSum += int(number[i]-'0')*2 - 9
			} else {
				checkSum += int(number[i]-'0') * 2
			}
		} else {
			checkSum += int(number[i] - '0')
		}
		digitCount += 1
	}
	return digitCount > 1 && checkSum%10 == 0
}
