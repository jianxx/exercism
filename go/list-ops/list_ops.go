// Package listops implement list operations:
// Foldr, Foldl, Filter, Length, Map, Reverse, Append, Concat
package listops

type IntList []int
type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

// Foldl (*given a function, a list, and initial accumulator, fold (reduce) each item into the accumulator from the left using `function(accumulator, item)`*);
func (l IntList) Foldl(f binFunc, acc int) int {
	for _, i := range l {
		acc = f(acc, i)
	}
	return acc
}

// Foldr : (*given a function, a list, and an initial accumulator, fold (reduce) each item into the accumulator from the right using `function(item, accumulator)`*);
func (l IntList) Foldr(f binFunc, acc int) int {
	for i := len(l) - 1; i >= 0; i-- {
		acc = f(l[i], acc)
	}
	return acc
}

// Filter (*given a predicate and a list, return the list of all items for which `predicate(item)` is True*);
func (l IntList) Filter(p predFunc) IntList {
	r := make(IntList, 0, len(l))
	for _, v := range l {
		if p(v) {
			r = append(r, v)
		}
	}
	return r
}

// Length (*given a list, return the total number of items within it*);
func (l IntList) Length() int {
	return len(l)
}

// Map (*given a function and a list, return the list of the results of applying `function(item)` on all items*);
func (l IntList) Map(u unaryFunc) IntList {
	r := make(IntList, len(l))
	for i, v := range l {
		r[i] = u(v)
	}
	return r
}

// Reverse (*given a list, return a list with all the original items, but in reversed order*);
func (l IntList) Reverse() IntList {
	r := make(IntList, len(l))
	for i, v := range l {
		r[len(l)-1-i] = v
	}
	return r
}

// Append (*given two lists, add all items in the second list to the end of the first list*);
func (l IntList) Append(m IntList) IntList {
	r := make(IntList, 0, len(l)+len(m))
	r = append(r, l...)
	r = append(r, m...)
	return r
}

// Concat (*given a series of lists, combine all items in all lists into one flattened list*);
func (l IntList) Concat(a []IntList) IntList {
	r := make(IntList, 0, len(l)*len(a))
	r = append(r, l...)
	for _, m := range a {
		r = append(r, m...)
	}
	return r
}
