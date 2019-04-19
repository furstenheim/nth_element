package MediansOfMedians


import (
	"sort"
	"github.com/furstenheim/nth_element/utils"
)

// Based on https://en.wikipedia.org/wiki/Median_of_medians
// Buckets. Sort a slice into buckets of given size. All elements from one bucket are smaller than any element  from the next one.
// elements at position i * bucketSize are guaranteed to be the (i * bucketSize) th smallest elements
// s := // some slice
// FloydRivest.Buckets(sort.Interface(s), 5)
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
	_ = _select(array, k, left, right)

}
func _select(array sort.Interface, k, left, right int) int {
	for true {
		if left == right {
			return left
		}
		pivotIndex := pivotFunc(array, left, right)
		pivotIndex = partition(array, k, left, right, pivotIndex)
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
// Compute approximation of median by computing the median of subgroups of size 5,
// Pushing them to front and then computing the median of that
func pivotFunc(array sort.Interface, left, right int) int {
	// for 5 or less elements we just get the median
	if right - left < 5 {
		return insertionSortMedian(array, left, right)
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
		median5 := insertionSortMedian(array, i, subRight)
		array.Swap(median5, left + (i - left) / 5) // actually j
		j++
	}
	// compute the median of the n/5 medians of five
	// log.Println(j, left, (j + left) / 2, array)
	mid := (right -left) / 10 + left + 1 // (j + left) / 2 not really

	return _select(array, mid, left, left + (right - left) / 5)
}

func insertionSortMedian (array sort.Interface, left, right int) int {
	i := left + 1
	for i <= right {
		j := i
		for j > left && array.Less(j, j - 1) {
			array.Swap(j - 1, j)
			j = j - 1
		}
		i++
	}
	return (left + right) / 2
}

func partition (array sort.Interface, k, left, right, pivotIndex int) int {
	// Move pivot to the end
	array.Swap(pivotIndex, right)
	storeIndex := left
	// Move all elements smaller than the pivot to the left of the pivot
	for i := left; i < right; i++ {
		// pivot value is stored in right position
		// so here we compare with pivot
		if array.Less(i, right) {
			array.Swap(storeIndex, i)
			storeIndex++
		}
	}
	// Now move all items equal to the pivot value after the smaller items
	storeIndexEq := storeIndex
	for i := storeIndexEq; i < right; i++ {
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
