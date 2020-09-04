package sorts

func merge(a []int, b []int) []int {
	var aCopy []int = make([]int, len(a))
	var bCopy []int = make([]int, len(b))
	copy(aCopy, a)
	copy(bCopy, b)
	if len(aCopy) == 0 {
		return bCopy
	} else if len(bCopy) == 0 {
		return aCopy
	} else if aCopy[0] <= bCopy[0] {
		return append(aCopy[:1], merge(aCopy[1:], bCopy)...)
	}
	return append(bCopy[:1], merge(aCopy, bCopy[1:])...)
}

func mergeSort(collection []int) []int {
	if len(collection) <= 1 {
		return collection
	}
	return merge(mergeSort(collection[:len(collection)/2]), mergeSort(collection[len(collection)/2:]))
}
