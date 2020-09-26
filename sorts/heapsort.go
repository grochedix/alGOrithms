package sorts

func heapify(arr []int, node int, n int) {
	k := node
	j := 2 * k
	for j <= n {
		if j < n && arr[j-1] < arr[j] {
			j++
		}
		if arr[k-1] < arr[j-1] {
			arr[k-1], arr[j-1] = arr[j-1], arr[k-1]
			k = j
			j = 2 * k
		} else {
			break
		}
	}
}

func heapSort(arr []int) []int {
	var a []int = make([]int, len(arr))
	copy(a, arr)
	for i := len(a) / 2; i > 0; i-- {
		heapify(a, i, len(a))
	}
	for i := len(a); i > 1; i-- {
		a[i-1], a[0] = a[0], a[i-1]
		heapify(a, 1, i-1)
	}
	return a
}
