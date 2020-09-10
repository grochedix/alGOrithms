package search

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSearch(t *testing.T) {
	collection := []int{0, 5, 4, 2, 3, 9, 1, 6, 0, 1}
	fmt.Println(linear(collection, 2))  // 3
	fmt.Println(linear(collection, 10)) //-1
	collection = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for i := range collection {
		fmt.Println(strconv.Itoa(i) + " " + strconv.Itoa(binarySearch(collection, i)))
	}
}
