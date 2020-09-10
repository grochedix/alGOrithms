package search

import "math"

func jumpSearch(arr []int, x int) int {
	n := float64(len(arr))
	step := math.Floor(math.Sqrt(n))
	for arr[int(math.Min(step, n-1))] < x {
		if step >= n {
			return -1
		}
		step += math.Floor(math.Sqrt(n))
	}
	for pos := int(math.Min(step, n-1)); arr[pos] >= x; pos-- {
		if arr[pos] == x && (pos == 0 || arr[pos-1] != x) {
			return pos
		}
		if pos == int(step-math.Floor(math.Sqrt(n))) {
			return -1
		}
	}
	return -1
}
