package sorts

func insertionSort(collection []int) []int {
	for i := 1; i < len(collection); i++ {
		for j := i; j > 0 && collection[j-1] > collection[j]; j-- {
			collection[j-1], collection[j] = collection[j], collection[j-1]
		}
	}
	return collection
}
