package sorts

func gnomeSort(collection []int) []int {
	for i := 1; i < len(collection); {
		if collection[i] < collection[i-1] {
			collection[i], collection[i-1] = collection[i-1], collection[i]
			if i > 1 {
				i--
			}
		} else {
			i++
		}
	}
	return collection
}
