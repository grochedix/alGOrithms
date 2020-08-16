package sorts

func bubbleSort(collection []int) []int {
	for i := 0; i < len(collection); i++ {
		swapped := false
		for j := 0; j < len(collection)-i-1; j++ {
			if collection[j+1] < collection[j] {
				collection[j], collection[j+1] = collection[j+1], collection[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return collection
}
