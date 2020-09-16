package raindrops

import "strconv"

// Convert can convert a number into a string that contains raindrop sounds
func Convert(number int) string {
	r := ""
	if number%3 == 0 {
		r += "Pling"
	}
	if number%5 == 0 {
		r += "Plang"
	}
	if number%7 == 0 {
		r += "Plong"
	}
	if r == "" {
		return strconv.Itoa(number)
	}
	return r
}
