package project_test

import (
	"testing"
	"fmt"
	"github.com/furstenheim/nth_element/FloydRivest"
	"github.com/furstenheim/nth_element/FullSort"
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
		{32},
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
	bucketSize := size / nBuckets
	b.Run("Floyd Rivest", func (b *testing.B) {
		for n := 0; n < b.N; n++ {
			inputArray = inputArray[0:0]
			for i := 0; i < size; i++ {
				inputArray = append(inputArray, ta.array[i])
			}
			FloydRivest.IntBuckets(nthElementUtils.IntSorter(inputArray), bucketSize)
		}
	})
	b.Run("Full Sort", func (b *testing.B) {
		for n := 0; n < b.N; n++ {
			inputArray = inputArray[0:0]
			for i := 0; i < size; i++ {
				inputArray = append(inputArray, ta.array[i])
			}
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
	var sortedArray = make([]int, MAX_SIZE_ARRAY)
	var invertedArray = make([]int, MAX_SIZE_ARRAY)
	var constantArray = make([]int, MAX_SIZE_ARRAY)
	var sawArray = make([]int, MAX_SIZE_ARRAY)
	for i, _ := range(sortedArray) {
		sortedArray[i] = i
		invertedArray[i] = len(invertedArray) - i
		sawArray[i] = i % 10
	}

	return []testArray {
		{
			"sorted",
			sortedArray,
		},
		{
			"inverted",
			invertedArray,
		},
		{
			"constant",
			constantArray,
		},
		{
			"sawLike",
			sawArray,
		},
	}
}

type testArray struct {
	name string
	array []int
}