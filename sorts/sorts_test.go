package sorts

import (
	"fmt"
	"testing"
)

func TestSorts(t *testing.T) {
	collection := []int{0, 5, 4, 2, 3, 9, 1, 6, 0, 1}
	fmt.Println(mergeSort(collection))
}
