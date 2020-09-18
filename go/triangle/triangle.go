// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Kind Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// NaT : not a triangle.
	NaT Kind = iota // not a triangle
	// Equ : equilateral
	Equ // equilateral
	// Iso : isosceles
	Iso // isosceles
	// Sca : scalene
	Sca // scalene
)

// KindFromSides determines if a triangle is equilateral, isosceles, or scalene.
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) || a <= 0 || b <= 0 || c <= 0 || a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a == b && a == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}
