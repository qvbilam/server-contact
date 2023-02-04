package utils

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	//arr := []int{1, 2, 3, 4, 5}
	arr := []int{5, 4, 3, 2, 1}
	//arr := []int{1, 3, 5, 6, 4}
	need := 1

	index := BinarySearch(arr, need)
	fmt.Println(index)
}
