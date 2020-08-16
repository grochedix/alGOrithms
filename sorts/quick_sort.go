package sorts

func quickSort(collection []int) []int {
	if len(collection) <= 1 {
		return collection
	}

	var lesser []int
	var greater []int
	var pivot int = collection[len(collection)-1]

	for i := 0; i < len(collection)-1; i++ {
		if pivot < collection[i] {
			greater = append(greater, collection[i])
		} else {
			lesser = append(lesser, collection[i])
		}
	}

	return append(append(quickSort(lesser), pivot), quickSort(greater)...)
}
