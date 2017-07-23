package binarysearch

import (
	"testing"
)

type inttestgroup struct {
	input SearchInt
	key   int
	index int
}

type stringtestgroup struct {
	input SearchString
	key   string
	index int
}

func testGroupSetupInt(t *testing.T) []inttestgroup {
	tests := []inttestgroup{
		{SearchInt{[]int{}}, 3, -1},
		{SearchInt{[]int{1, 2, 3, 4}}, 3, 2},
		{SearchInt{[]int{1, 2, 3, 4}}, 7, -1},
		{SearchInt{[]int{8, 9, 10, 11}}, 7, -1},
		{SearchInt{[]int{8, 10, 9, 11}}, 9, -1}, // incorrectly sorted
		{SearchInt{[]int{3, 5, 8, 10, 12, 15, 18, 20, 20, 50, 60}}, 50, 9},
		{SearchInt{[]int{1, 5, 8, 12, 13}}, 8, 2},
		{SearchInt{[]int{1, 5, 8, 12, 13}}, 1, 0},
		{SearchInt{[]int{1, 5, 8, 12, 13}}, 23, -1},
		{SearchInt{[]int{1, 5, 8, 12, 13}}, 1, 0},
		{SearchInt{[]int{1, 5, 8, 12, 13}}, 11, -1},
	}
	var bigarray []int
	for i := 1; i <= 1000; i++ {
		bigarray = append(bigarray, i)
	}
	bigtest := inttestgroup{SearchInt{bigarray}, 150, 149}
	tests = append(tests, bigtest)

	return tests
}

func TestBinarySearchInt(t *testing.T) {
	for _, group := range testGroupSetupInt(t) {
		ret := group.input.BinarySearch(0, len(group.input.List)-1, group.key)
		if ret != group.index {
			t.Error(
				"For", group.input,
				"expected", group.index,
				"got", ret,
			)
		}
	}

}

func testGroupSetupString(t *testing.T) []stringtestgroup {
	tests := []stringtestgroup{
		{SearchString{[]string{}}, "blah", -1},
		{SearchString{[]string{"one", "two", "three", "four"}}, "three", -1},
		{SearchString{[]string{"one", "three", "two", "four"}}, "four", -1},
		{SearchString{[]string{"four", "one", "three", "two"}}, "four", 0},
		{SearchString{[]string{"four", "one", "three", "two"}}, "one", 1},
		{SearchString{[]string{"four", "one", "three", "two"}}, "three", 2},
		{SearchString{[]string{"four", "one", "three", "two"}}, "two", 3},
		{SearchString{[]string{"a", "n", "w", "z"}}, "n", 1},
		{SearchString{[]string{"a", "n", "w", "z"}}, "z", 3},
	}

	return tests
}

func TestBinarySearchString(t *testing.T) {
	for _, group := range testGroupSetupString(t) {
		ret := group.input.BinarySearch(0, len(group.input.List)-1, group.key)
		if ret != group.index {
			t.Error(
				"For", group.input,
				"expected", group.index,
				"got", ret,
			)
		}
	}

}
