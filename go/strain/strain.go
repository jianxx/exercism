package strain

// Ints ...
type Ints []int

// Lists ...
type Lists [][]int

// Strings ...
type Strings []string

// Keep returns a new collection containing those elements where the predicate is true
func (i Ints) Keep(f func(int) bool) Ints {
	r := Ints{}
	for _, e := range i {
		if f(e) {
			r = append(r, e)
		}
	}
	if len(r) == 0 {
		return nil
	}
	return r
}

// Discard returns a new collection containing those elements where the predicate is false
func (i Ints) Discard(f func(int) bool) Ints {
	return i.Keep(func(n int) bool { return !f(n) })
}

// Keep returns a new collection containing those elements where the predicate is true
func (l Lists) Keep(f func([]int) bool) Lists {
	r := Lists{}
	for _, e := range l {
		if f(e) {
			r = append(r, e)
		}
	}
	if len(r) == 0 {
		return nil
	}
	return r
}

// Keep returns a new collection containing those elements where the predicate is true
func (ss Strings) Keep(f func(string) bool) Strings {
	r := Strings{}
	for _, e := range ss {
		if f(e) {
			r = append(r, e)
		}
	}
	if len(r) == 0 {
		return nil
	}
	return r
}
