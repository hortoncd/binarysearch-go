package binarysearch

type SearchInt struct {
	List []int
	Key  int
}

type SearchString struct {
	List []string
	Key  string
}

type Searchable interface {
	binarySearch() int
}

// BinarySearch takes and interface and calls the search for the specific
// data type of the arg, in order to support strings with ints
func BinarySearch(s Searchable) int {
	return s.binarySearch()
}

// binarySearch handles calling the int-specific BinarySearch with args
func (s SearchInt) binarySearch() int {
	return s.doBinarySearch(0, len(s.List)-1, s.Key)
}

// BinarySearch implements the binary search algorithm on an array of integers
// assuming that the input array is already sorted
func (s SearchInt) doBinarySearch(low int, high int, key int) int {

	m := mid(low, high)
	if m < 0 {
		return -1
	}

	if s.List[m] == key {
		return m
	}

	if key < s.List[m] {
		return s.doBinarySearch(low, m-1, key)
	} else {
		return s.doBinarySearch(m+1, high, key)
	}

}

// binarySearch handles calling the string-specific BinarySearch with args
func (s SearchString) binarySearch() int {
	return s.doBinarySearch(0, len(s.List)-1, s.Key)
}

// BinarySearch implements the binary search algorithm on an array of strings
// assuming that the input array is already sorted
func (s SearchString) doBinarySearch(low int, high int, key string) int {

	m := mid(low, high)
	if m < 0 {
		return -1
	}

	if s.List[m] == key {
		return m
	}

	if key < s.List[m] {
		return s.doBinarySearch(low, m-1, key)
	} else {
		return s.doBinarySearch(m+1, high, key)
	}

}

// mid tests high/low for validity and returns mid point
func mid(low int, high int) int {

	if high < low {
		return -1
	}

	return low + ((high - low) / 2)

}
