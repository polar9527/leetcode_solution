package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	array := []int{3, 2, 7, 4, 9, 7, 8, 6, 7, 11, 0, 4, 5}
	bubbleSort(array)
	t.Log(array)
}
