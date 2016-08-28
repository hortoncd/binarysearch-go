package binarysearch

// BinarySearch implements the binary search algorithm on an array of integers
// assuming that the input array is already sorted
func BinarySearch(input *[]int, low int, high int, key int) int {

	if high < low {
		return -1
	}

	mid := low + ((high - low) / 2)

	if (*input)[mid] == key {
		return mid
	}

	if key < (*input)[mid] {
		return BinarySearch(input, low, mid-1, key)
	} else {
		return BinarySearch(input, mid+1, high, key)
	}
}
