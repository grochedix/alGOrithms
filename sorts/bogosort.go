package sorts

import "math/rand"

func isOrdered(collection []int) bool {
	for i := 0; i < len(collection)-1; i++ {
		if collection[i] > collection[i+1] {
			return false
		}
	}
	return true
}

func bogosort(collection []int) []int {
	if collection == nil {
		return collection
	}
	for {
		var sorted = make([]int, 0, len(collection))
		var perm []int = rand.Perm(len(collection))
		for _, v := range perm {
			sorted = append(sorted, collection[v])
		}
		if isOrdered(sorted) {
			collection = sorted
			break
		}
	}
	return collection
}
