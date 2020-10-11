package collatzconjecture

import "fmt"

func CollatzConjecture(input int) (int, error) {
	if input <= 0 {
		return 0, fmt.Errorf("Invalid input: %d", input)
	}
	if input == 1 {
		return 0, nil
	}

	var steps = 0
	for input != 1 {
		if input%2 == 0 {
			input /= 2
		} else {
			input = 3*input + 1
		}
		steps += 1
	}
	return steps, nil
}
