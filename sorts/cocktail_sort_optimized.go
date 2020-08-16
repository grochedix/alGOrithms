package sorts

func cocktailSortOptimized(collection []int) []int {
	var swapped bool = true
	start := 0
	end := len(collection) - 2
	for swapped {
		swapped = false
		for j := start; j <= end; j++ {
			if collection[j+1] < collection[j] {
				collection[j], collection[j+1] = collection[j+1], collection[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		end--
		for j := end; j >= start; j-- {
			if collection[j+1] < collection[j] {
				collection[j], collection[j+1] = collection[j+1], collection[j]
				swapped = true
			}
		}
		start++
	}
	return collection
}
