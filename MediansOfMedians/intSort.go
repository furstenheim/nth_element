package MediansOfMedians


import (
	"github.com/furstenheim/nth_element/utils"
)

// Based on https://en.wikipedia.org/wiki/Median_of_medians
// IntBuckets. Sort a slice into buckets of given size. All elements from one bucket are smaller than any element  from the next one.
// elements at position i * bucketSize are guaranteed to be the (i * bucketSize) th smallest elements
// s := // some slice
// FloydRivest.IntBuckets(nthElementUtils.IntSorter(s), 5)
// max(s[0:5]) < min(s[5:10])
// max(s[10: 15]) < min(s[15:20])
// ...
func IntBuckets(slice nthElementUtils.IntSorter, bucketSize int) {
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
		IntSelect(slice, mid, left, right)
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
func IntSelect(array nthElementUtils.IntSorter, k, left, right int) {
	_ = int__select(array, k, left, right)

}
func int__select(array nthElementUtils.IntSorter, k, left, right int) int {
	for true {
		if left == right {
			return left
		}
		pivotIndex := int_pivotFunc(array, left, right)
		pivotIndex = int_partition(array, k, left, right, pivotIndex)
		if k == pivotIndex {
			return k
		} else if k < pivotIndex {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}
	return -1
}

func int_pivotFunc(array nthElementUtils.IntSorter, left, right int) int {
	// for 5 or less elements we just get the median
	if right - left < 5 {
		return int_insertionSortMedian(array, left, right)
	}
	j := left
	// move medians of five element subgroups to the first n/5 positions
	for i := left; i < right; i += 5 {
		// get the median of the ith five-element subgroup
		// TODO we can probably split this outside of the for
		subRight := i + 4
		if subRight > right {
			subRight = right
		}
		median5 := int_insertionSortMedian(array, i, subRight)
		array.Swap(median5, j)
		j++
	}
	// compute the median of the n/5 medians of five
	mid := (j + left) / 2
	return int__select(array, mid, left, j)
}

func int_insertionSortMedian (array nthElementUtils.IntSorter, left, right int) int {
	i := left + 1
	for i <= right {
		j := i
		for j > left && array.Less(j - 1, j) {
			array.Swap(j - 1, j)
			j = j - 1
		}
		i++
	}
	return (left + right) / 2
}

func int_partition (array nthElementUtils.IntSorter, k, left, right, pivotIndex int) int {
	// Move pivot to the end
	array.Swap(pivotIndex, right)
	storeIndex := left
	// Move all elements smaller than the pivot to the left of the pivot
	for i := left; i < right - 1; i++ {
		// pivot value is stored in right position
		// so here we compare with pivot
		if array.Less(i, right) {
			array.Swap(storeIndex, i)
			storeIndex++
		}
	}
	// Now move all items equal to the pivot value after the smaller items
	storeIndexEq := storeIndex
	for i := storeIndexEq; i < right - 1; i++ {
		if !array.Less(i, right) && !array.Less(right, i) {
			array.Swap(storeIndexEq, i)
			storeIndexEq++
		}
	}
	array.Swap(right, storeIndexEq)
	// At this point pivotValue is in storeIndexEq, to the left of storeIndex everything is
	// smaller. And to the left of storeIndexEq everything is smaller or equal

	// Return location of pivot considering location n
	if k < storeIndex {
		return storeIndex // k is in the group of smaller items
	}
	if k <= storeIndexEq {
		return k // k is in the group equal to the pivot
	}
	return storeIndexEq // k is in the group of larger items
}
