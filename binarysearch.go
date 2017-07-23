package binarysearch

type SearchInt struct {
	List []int
}

type SearchString struct {
	List []string
}

// BinarySearch implements the binary search algorithm on an array of integers
// assuming that the input array is already sorted
func (s SearchInt) BinarySearch(low int, high int, key int) int {

	m := mid(low, high)
	if m < 0 {
		return -1
	}

	if s.List[m] == key {
		return m
	}

	if key < s.List[m] {
		return s.BinarySearch(low, m-1, key)
	} else {
		return s.BinarySearch(m+1, high, key)
	}

}

// BinarySearch implements the binary search algorithm on an array of strings
// assuming that the input array is already sorted
func (s SearchString) BinarySearch(low int, high int, key string) int {

	m := mid(low, high)
	if m < 0 {
		return -1
	}

	if s.List[m] == key {
		return m
	}

	if key < s.List[m] {
		return s.BinarySearch(low, m-1, key)
	} else {
		return s.BinarySearch(m+1, high, key)
	}

}

// mid tests high/low for validity and returns mid point
func mid(low int, high int) int {

	if high < low {
		return -1
	}

	return low + ((high - low) / 2)

}
