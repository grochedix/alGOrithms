package sorts

func stoogeSort(collection []int, start int, end int) []int {
	if collection[start] > collection[end-1] {
		collection[end-1], collection[start] = collection[start], collection[end-1]
	}
	if (end - start) > 2 {
		t := (end - start) / 3
		stoogeSort(collection, start, end-t)
		stoogeSort(collection, start+t, end)
		stoogeSort(collection, start, end-t)
	}
	return collection
}
