package in

// Convenience functions for golangs silly lack of set operations on slices.
// The lists to search in will get sorted as a side-effect
// This is fast  for small slices, if you have a larger slice you are probably
// better off using a map.

// Strings checks existence of a string in a string slice
func Strings(haystack []string, needle string) bool {
	for _, elem := range haystack {
		if elem == needle {
			return true
		}
	}
	return false
}

// Ints checks existence of an int in an int slice
func Ints(haystack []int, needle int) bool {
	for _, elem := range haystack {
		if elem == needle {
			return true
		}
	}
	return false
}

// StringSubset checks if needles is a subset of haystack
func StringsSubset(haystack []string, needles []string) bool {
	set := make(map[string]bool)
	for _, value := range haystack {
		set[value] = true
	}

	for _, value := range needles {
		if _, found := set[value]; !found {
			return false
		}
	}
	return true
}

// IntSubset checks if needles is a subset of haystack
func IntsSubset(haystack []int, needles []int) bool {
	set := make(map[int]bool)
	for _, value := range haystack {
		set[value] = true
	}

	for _, value := range needles {
		if _, found := set[value]; !found {
			return false
		}
	}
	return true
}
