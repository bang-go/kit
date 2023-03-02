package bslice

import (
	"github.com/bang-go/kit/constraint"
	"sort"
)

// ContainValue reports whether v is present in s.
func ContainValue[E comparable](s []E, v E) bool {
	return Index(s, v) >= 0
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func Index[E comparable](s []E, v E) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}

func Sort[T constraint.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}
