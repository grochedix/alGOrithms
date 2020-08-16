package sorts

func bubbleSort(collection []int) []int {
	var swapped bool = true
	for i := 1; i < len(collection); i++ {
		swapped = false
		for j := 1; j < len(collection); {
			if collection[j-1] > collection[j] {
				collection[j], collection[j-1] = collection[j-1], collection[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return collection
}
