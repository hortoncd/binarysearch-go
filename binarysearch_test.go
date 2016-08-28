package binarysearch

import (
	"testing"
)

type testgroup struct {
	input []int
	key   int
	index int
}

func testGroupSetup(t *testing.T) []testgroup {
	tests := []testgroup{
		{[]int{}, 3, -1},
		{[]int{1, 2, 3, 4}, 3, 2},
		{[]int{1, 2, 3, 4}, 7, -1},
		{[]int{8, 9, 10, 11}, 7, -1},
		{[]int{8, 10, 9, 11}, 9, -1}, // incorrectly sorted
		{[]int{3, 5, 8, 10, 12, 15, 18, 20, 20, 50, 60}, 50, 9},
		{[]int{1, 5, 8, 12, 13}, 8, 2},
		{[]int{1, 5, 8, 12, 13}, 1, 0},
		{[]int{1, 5, 8, 12, 13}, 23, -1},
		{[]int{1, 5, 8, 12, 13}, 1, 0},
		{[]int{1, 5, 8, 12, 13}, 11, -1},
	}
	var bigarray []int
	for i := 1; i <= 1000; i++ {
		bigarray = append(bigarray, i)
	}
	bigtest := testgroup{bigarray, 150, 149}
	tests = append(tests, bigtest)

	return tests
}

func TestBinarySearch(t *testing.T) {
	for _, group := range testGroupSetup(t) {
		ret := BinarySearch(&group.input, 0, len(group.input)-1, group.key)
		if ret != group.index {
			t.Error(
				"For", group.input,
				"expected", group.index,
				"got", ret,
			)
		}
	}
}
