package project_test

import (
	"testing"
	"fmt"
	"github.com/furstenheim/nth_element/FloydRivest"
	"github.com/furstenheim/nth_element/FullSort"
	"github.com/furstenheim/nth_element/QuickSelect"
	"github.com/furstenheim/nth_element/MediansOfMedians"
	"github.com/furstenheim/nth_element/utils"

)
const MAX_SIZE_ARRAY = 1000000


func BenchmarkBuckets (b *testing.B) {
	testArrays := getTestArrays()
	testSizes:= []struct{
		n int
	} {
		{100},
		{1000},
		{10000},
		{100000},
		{MAX_SIZE_ARRAY},
	}
	buckets := []struct{
		n int
	} {
		{2},
		{9},
		//{32},
	}
	var inputArray = make([]int, MAX_SIZE_ARRAY)

	for _, bucket := range(buckets) {
		for _, ts := range(testSizes) {
			b.Run(fmt.Sprintf("Size %d/Buckets %d", ts.n, bucket.n), func (b *testing.B) {
				for _, t := range(testArrays) {
					b.Run(t.name, func (b * testing.B) {
						benchArray(b, ts.n, bucket.n, t, inputArray)
					})
				}

			})
		}
	}
}

func benchArray (b * testing.B, size, nBuckets int, ta testArray, inputArray []int) {
	benchCases := []struct {
		name string
		maxSize int // avoid O(n2) in quickselect worst case
		algorithm func (nthElementUtils.IntSorter, int)
	} {
		{
			"FloydRivest",
			-1,
			FloydRivest.IntBuckets,
		},
		{
			"QuickSelect",
			1000000,
			QuickSelect.IntBuckets,
		},
		{
			"MediansOfMedians",
			-1,
			MediansOfMedians.IntBuckets,
		},
	}
	bucketSize := size / nBuckets
	for _, bc := range(benchCases) {
		if bc.maxSize == -1 || size < bc.maxSize {
			b.Run(bc.name, func (b * testing.B) {
				for n := 0; n < b.N; n++ {
					inputArray = inputArray[0:size]
					ta.setUpFunction(inputArray)
					bc.algorithm(nthElementUtils.IntSorter(inputArray), bucketSize)
				}
			})
		}

	}
	b.Run("Full Sort", func (b *testing.B) {
		for n := 0; n < b.N; n++ {
			inputArray = inputArray[0:0]
			inputArray = inputArray[0:size]
			ta.setUpFunction(inputArray)
			FullSort.Sort(nthElementUtils.IntSorter(inputArray))
		}
	})

}


/*
func BenchmarkSort200kSize5(b *testing.B) {
	for n:= 0; n < b.N; n++ {
		copyArray := append([]int{}, benchArray200k...)
		sort.Sort(sorter(copyArray))
	}
}

func BenchmarkBuckets200knBuckets5(b *testing.B) {
	for n:= 0; n < b.N; n++ {
		copyArray := append([]int{}, benchArray200k...)
		Buckets(sorter(copyArray), len(copyArray) / 5)
	}
}

func BenchmarkBuckets200knBuckets16(b *testing.B) {
	for n:= 0; n < b.N; n++ {
		copyArray := append([]int{}, benchArray200k...)
		Buckets(sorter(copyArray), len(copyArray) / 16)
	}
}

func BenchmarkBuckets200knBuckets32(b *testing.B) {
	for n:= 0; n < b.N; n++ {
		copyArray := append([]int{}, benchArray200k...)
		Buckets(sorter(copyArray), len(copyArray) / 32)
	}
}*/

func getTestArrays ()  []testArray {
	return []testArray {
		{
			"sorted",
			func (array []int) {
				for i, _ := range(array) {
					array[i] = i
				}
			},
		},
		{
			"inverted",
			func (array []int) {
				for i, _ := range(array) {
					array[i] = len(array) - i
				}
			},
		},
		{
			"constant",
			func (array []int) {
				for i, _ := range(array) {
					array[i] = 0
				}
			},
		},
		{
			"sawLike",
			func (array []int) {
				for i, _ := range(array) {
					array[i] = i % (len(array) / 100)
				}
			},
		},
		{
			"pyramid",
			func (array []int) {
				for i, _ := range(array) {
					if i < len(array) / 2 {
						array[i] = i
					} else {
						array[i] = len(array) - i
					}
				}
			},
		},
	}
}

type testArray struct {
	name string
	setUpFunction func (array []int)
}