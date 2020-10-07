package protein

import (
	"fmt"
)

// ErrStop is an error implementation that includes a time and message.
type ErrStop struct {
}

// ErrInvalidBase is an error implementation that includes a time and message.
type ErrInvalidBase struct {
}

func (e ErrStop) Error() string {
	return fmt.Sprintf("ErrStop")
}

func (e ErrInvalidBase) Error() string {
	return fmt.Sprintf("InvalidBase")
}

var translater = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

// FromCodon translates codon sequences(string) into proteins([]string)
func FromCodon(codon string) (string, error) {
	protein, exists := translater[codon]
	if !exists {
		return "", ErrInvalidBase{}
	}
	if protein == "STOP" {
		return "", ErrStop{}
	}
	return protein, nil
}

// FromRNA translates RNA sequences(string) into proteins([]string)
func FromRNA(input string) ([]string, error) {
	var proteins []string
	for i := 0; i < len(input); i += 3 {
		protein, err := FromCodon(input[i : i+3])
		if err != nil {
			switch err.(type) {
			case ErrStop:
				return proteins, nil
			default:
				return proteins, err
			}
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}
