package search

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	collection := []int{0, 5, 4, 2, 3, 9, 1, 6, 0, 1}
	fmt.Println(linear(collection, 2))  // 3
	fmt.Println(linear(collection, 10)) //-1
}
