package hamming

import (
	"fmt"
)

// Distance returns the Hamming Distance between two strings.
func Distance(a, b string) (int, error) {
	diff := 0
	if len(a) != len(b) {
		return 0, fmt.Errorf("a and b are difference size")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff, nil
}
