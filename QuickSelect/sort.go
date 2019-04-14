// Based on https://en.wikipedia.org/wiki/Quickselect
package QuickSelect

import (
	"sort"
	"github.com/furstenheim/nth_element/utils"
)

// Buckets. Sort a slice into buckets of given size. All elements from one bucket are smaller than any element  from the next one.
// elements at position i * bucketSize are guaranteed to be the (i * bucketSize) th smallest elements
// s := // some slice
// QuickSelect.Buckets(sort.Interface(s), 5)
// s is now sorted into buckets of size 5
// max(s[0:5]) < min(s[5:10])
// max(s[10: 15]) < min(s[15:20])
// ...
func Buckets(slice sort.Interface, bucketSize int) {
	left := 0
	right := slice.Len() - 1
	s := nthElementUtils.Stack([]int{left, right})
	var mid int
	for len(s) > 0 {
		s, right = s.Pop()
		s, left = s.Pop()
		if (right - left <= bucketSize) {
			continue
		}
		// + bucketSize - 1 is to do math ceil
		mid = left + ((right - left + bucketSize - 1) / bucketSize / 2) * bucketSize
		Select(slice, mid, left, right)
		s = s.Push(left)
		s = s.Push(mid)
		s = s.Push(mid)
		s = s.Push(right)
	}
}

// left is the left index for the interval
// right is the right index for the interval
// k is the desired index value, where array[k] is the k+1 smallest element
// when left = 0
func Select(array sort.Interface, k, left, right int) {
	for true {
		if left == right {
			return
		}
		pivotIndex := (left + right) >> 1 // Use middle point as pivot. This could probably use random index
		pivotIndex = partition(array, pivotIndex, left, right)
		if k == pivotIndex {
			return
		} else if k < pivotIndex {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}

	}
}

// Partition values into less than array[pivot] and greater than array[pivot]
func partition (array sort.Interface, pivot, left, right int) int {
	array.Swap(pivot, right) // Move pivot to end
	storeIndex := left
	for i := left; i < right; i++ {
		if array.Less(i, right) { // Compare to pivot value
			array.Swap(storeIndex, i)
			storeIndex++
		}
	}
	array.Swap(right, storeIndex) // Return pivot Index to position
	return storeIndex

}
