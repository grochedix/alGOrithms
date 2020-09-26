package permutations

func heaps(arr []int) [][]int {
	var perm []int = make([]int, len(arr))
	copy(perm, arr)
	var res [][]int

	var generate func(a []int, k int)
	generate = func(a []int, k int) {
		if k == 1 {
			var z []int = make([]int, len(a))
			copy(z, a) // to avoid pointer effect
			res = append(res, z)
			return
		}
		generate(a, k-1)
		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				a[i], a[k-1] = a[k-1], a[i]
			} else {
				a[0], a[k-1] = a[k-1], a[0]
			}
			generate(a, k-1)
		}
	}

	generate(perm, len(perm))
	return res
}
