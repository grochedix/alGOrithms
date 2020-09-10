package search

func binarySearch(ar []int, x int) int {
	if ar[0] > x || x > ar[len(ar)-1] {
		return -1
	}
	var mid int = len(ar) / 2
	if ar[mid] == x {
		return mid
	} else if ar[mid] < x {
		return mid + 1 + binarySearch(ar[mid+1:], x)
	}
	return binarySearch(ar[:mid], x)
}

func leftBinarySearch(ar []int, x int) int {
	if ar[0] > x || x > ar[len(ar)-1] {
		return -1
	}
	var mid int = len(ar) / 2
	if ar[mid] == x && (mid == 0 || ar[mid-1] != x) {
		return mid
	} else if ar[mid] < x {
		return mid + 1 + binarySearch(ar[mid+1:], x)
	}
	return binarySearch(ar[:mid], x)
}

func rightBinarySearch(ar []int, x int) int {
	if ar[0] > x || x > ar[len(ar)-1] {
		return -1
	}
	var mid int = len(ar) / 2
	if ar[mid] == x && (mid == len(ar)-1 || ar[mid+1] != x) {
		return mid
	} else if ar[mid] <= x {
		return mid + 1 + binarySearch(ar[mid+1:], x)
	}
	return binarySearch(ar[:mid], x)
}
