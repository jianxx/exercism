package strand

import "bytes"

// ToRNA return its RNA complement for the input dna
func ToRNA(dna string) string {
	var b bytes.Buffer
	for _, c := range dna {
		switch c {
		case 'G':
			b.WriteString("C")
		case 'C':
			b.WriteString("G")
		case 'T':
			b.WriteString("A")
		case 'A':
			b.WriteString("U")
		}
	}
	return b.String()
}
