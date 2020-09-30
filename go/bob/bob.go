// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.Join(strings.Fields(remark), "")
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}
	var isQuestion = strings.HasSuffix(remark, "?")
	var isYell = strings.IndexFunc(remark, unicode.IsLetter) >= 0 && strings.ToUpper(remark) == remark
	if isQuestion && isYell {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion && !isYell {
		return "Sure."
	} else if !isQuestion && isYell {
		return "Whoa, chill out!"
	}
	return "Whatever."
}
