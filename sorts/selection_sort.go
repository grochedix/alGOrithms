package sorts

import "fmt"

func minSlice(collection []int) (int, int) {
	if collection == nil {
		panic(fmt.Sprint("nil collection"))
	}
	var min int = collection[0]
	var index int = 0
	for i := 1; i < len(collection); i++ {
		if collection[i] < min {
			min = collection[i]
			index = i
		}
	}
	return min, index
}

func selectionSort(collection []int) []int {
	for i := 0; i < len(collection)-1; i++ {
		temp := collection[i:]
		_, index := minSlice(temp)
		collection[i], collection[index+i] = collection[index+i], collection[i]
	}
	return collection
}
