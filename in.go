package in

import "sort"

// Convenience functions for golangs silly lack of set operations on slices.
// The lists to search in will get sorted as a side-effect
// This is fast  for small slices, if you have a larger slice you are probably
// better off using a map.

// Strings check existence of a string in a string slice
func Strs(haystack []string, needle string) bool {
	sort.Strings(haystack)
	i := sort.SearchStrings(haystack, needle)
	return i < len(haystack) && haystack[i] == needle
}

// Ints checks existence of an int in an int slice
func Ints(haystack []int, needle int) bool {
	sort.Ints(haystack)
	i := sort.SearchInts(haystack, needle)
	return i < len(haystack) && haystack[i] == needle
}
