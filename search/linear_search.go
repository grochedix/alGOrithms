package search

func linear(ar []int, x int) int {
	for ind, val := range ar {
		if val == x {
			return ind
		}
	}
	return -1
}
