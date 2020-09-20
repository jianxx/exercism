package accumulate
<<<<<<< HEAD

// Accumulate will perform an operation on each element of a given collection
func Accumulate(input []string, operation func(string) string) []string {
	var r = []string{}
	for _, e := range input {
		r = append(r, operation(e))
	}
	return r
}
=======
>>>>>>> f81b578... Add all to practice branch
