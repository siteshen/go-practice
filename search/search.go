package search

// binary search 1: low ≤ mid < high
func BinarySearch1(list []int, v int) int {
	low, high := 0, len(list)
	for low < high {
		mid := (high-low)/2 + low
		// low ≤ mid < high
		if list[mid] < v {
			low = mid + 1
		} else if list[mid] > v {
			high = mid
		} else {
			return mid
		}
	}
	return -1
}

// binary search 1 with counting: low ≤ mid < high
func BinarySearchWithCounting1(list []int, v int) (pos, count int) {
	low, high := 0, len(list)
	for low < high {
		mid := (high-low)/2 + low
		count += 1
		// low ≤ mid < high
		if list[mid] < v {
			low = mid + 1
		} else if list[mid] > v {
			high = mid
		} else {
			return mid, count
		}
	}
	return -1, count
}

// binary search 2: low ≤ mid ≤ high
func BinarySearch2(list []int, v int) int {
	low, high := 0, len(list)-1
	for low <= high {
		mid := (low + high) / 2
		// low ≤ mid ≤ high
		if list[mid] < v {
			low = mid + 1
		} else if list[mid] > v {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// binary search 2 with counting: low ≤ mid ≤ high
func BinarySearchWithCounting2(list []int, v int) (pos, count int) {
	low, high := 0, len(list)-1
	for low <= high {
		mid := (low + high) / 2
		count += 1
		// low ≤ mid ≤ high
		if list[mid] < v {
			low = mid + 1
		} else if list[mid] > v {
			high = mid - 1
		} else {
			return mid, count
		}
	}
	return -1, count
}

// friend of test
var BinarySearch = BinarySearch1
var BinarySearchWithCounting = BinarySearchWithCounting1
