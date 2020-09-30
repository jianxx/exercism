// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"bytes"
	"log"
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	re, err := regexp.Compile("[^a-zA-Z0-9 -]+")
	if err != nil {
		log.Fatal(err)
	}
	s = re.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, "-", " ")
	words := strings.Fields(s)

	var b bytes.Buffer
	for i := range words {
		b.WriteString(string(words[i][0]))
	}
	return strings.ToUpper(b.String())
}
