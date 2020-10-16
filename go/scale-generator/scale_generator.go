package scale

import "strings"

func Scale(tonic, interval string) (scale []string) {
	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#":
		scale = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		scale = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	}

	tonic = strings.Title(tonic)
	for i, e := range scale {
		if e == tonic {
			scale = append(scale[i:], scale[:i]...)
			break
		}
	}

	if interval == "" {
		return scale
	}

	var partialScale []string
	stepSize := map[string]int{"m": 1, "M": 2, "A": 3}
	i := 0
	for _, diff := range strings.Split(interval, "") {
		if step, ok := stepSize[diff]; ok {
			partialScale = append(partialScale, scale[i%len(scale)])
			i += step
		}
	}

	return partialScale
}
