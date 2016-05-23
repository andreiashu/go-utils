package search

import (
	"sort"
)

// StringSlicesEqual compares two slices. The comparison case sensitive
// The order of the elements in each slice is not important
func StringSlicesEqual(a, b []string) bool {
	for _, av := range a {
		if !StringInSlice(av, b) {
			return false
		}
	}

	for _, bv := range b {
		if !StringInSlice(bv, a) {
			return false
		}
	}

	return true
}

// StringInSlice searches for a string in a slice
// The search is case sensitive
func StringInSlice(s string, slice []string) bool {
	sort.Strings(slice)
	i := sort.Search(len(slice), func(i int) bool { return slice[i] >= s })
	if i < len(slice) && slice[i] == s {
		return true
	}

	return false
}
