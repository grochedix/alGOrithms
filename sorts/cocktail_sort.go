package sorts

func cocktailSort(collection []int) []int {
	var swapped bool = true
	for swapped {
		swapped = false
		for j := 0; j < len(collection)-1; j++ {
			if collection[j+1] < collection[j] {
				collection[j], collection[j+1] = collection[j+1], collection[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		for j := len(collection) - 2; j >= 0; j-- {
			if collection[j+1] < collection[j] {
				collection[j], collection[j+1] = collection[j+1], collection[j]
				swapped = true
			}
		}
	}
	return collection
}
