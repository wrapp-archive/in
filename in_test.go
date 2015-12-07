package in

import (
	"testing"
)

func TestStrings(t *testing.T) {
	result1 := Strings([]string{"11", "22"}, "22")
	if result1 != true {
		t.Error("Expected to be true, got ", result1)
	}

	result2 := Strings([]string{"11", "22"}, "33")
	if result2 != false {
		t.Error("Expected to be false, got ", result2)
	}
}

func TestInts(t *testing.T) {
	result1 := Ints([]int{11, 22}, 22)
	if result1 != true {
		t.Error("Expected to be true, got ", result1)
	}

	result2 := Ints([]int{11, 22}, 33)
	if result2 != false {
		t.Error("Expected to be false, got ", result2)
	}
}

func TestStringsSubset(t *testing.T) {
	result1 := StringsSubset([]string{"11", "22"}, []string{"22"})
	if result1 != true {
		t.Error("Expected to be true, got ", result1)
	}

	result2 := StringsSubset([]string{"11", "22"}, []string{"33"})
	if result2 != false {
		t.Error("Expected to be false, got ", result2)
	}
}

func TestIntsSubset(t *testing.T) {
	result1 := IntsSubset([]int{11, 22}, []int{22})
	if result1 != true {
		t.Error("Expected to be true, got ", result1)
	}

	result2 := IntsSubset([]int{11, 22}, []int{33})
	if result2 != false {
		t.Error("Expected to be false, got ", result2)
	}
}

func TestStringsIntersection(t *testing.T) {
	result1 := StringsIntersection([]string{"2", "5"}, []string{"1", "4", "5"})
	if result1 != true {
		t.Error("Expected to be true, got ", result1)
	}

	result2 := StringsIntersection([]string{"2", "5"}, []string{"1", "4", "6"})
	if result2 != false {
		t.Error("Expected to be false, got ", result2)
	}

}

func TestIntsIntersection(t *testing.T) {
	result1 := IntsIntersection([]int{2, 5}, []int{1, 4, 5})
	if result1 != true {
		t.Error("Expected to be true, got ", result1)
	}

	result2 := IntsIntersection([]int{2, 5}, []int{1, 4, 6})
	if result2 != false {
		t.Error("Expected to be false, got ", result2)
	}

}
