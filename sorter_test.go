package project_test

import (
	"testing"
	"sort"
	"reflect"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"github.com/furstenheim/nth_element/FloydRivest"
	"github.com/furstenheim/nth_element/MediansOfMedians"
)
//go:generate ./generate.sh

const sizeBench = 200000
var benchArray200k = make([]int, sizeBench)
func init () {
	for i := 0; i < sizeBench; i++ {
		benchArray200k[i] = rand.Int()
	}
}

func TestSelect_ConcreteExamples (t * testing.T) {
	testCases := []struct {
		name      string
		algorithm func (array sort.Interface, k, left, right int)
		examples  []struct {
			index int
			initial, expected []int
		}
	} {
		{
			"FloydRivest",
			FloydRivest.Select,
			[]struct {
				index int
				initial, expected []int
			} {
				{
					8,
					[]int{65, 28, 59, 33, 21, 56, 22, 95, 50, 12, 90, 53, 28, 77, 39},
					[]int{39, 28, 28, 33, 21, 12, 22, 50, 53, 56, 59, 65, 90, 77, 95},
				},

				{
					3,
					[]int{22, 33, 12, 95, 65, 28, 28, 77, 39, 21, 59, 50, 53, 56, 90},
					[]int{12, 21, 22, 28, 28, 33, 39, 50, 53, 56, 59, 65, 77, 90, 95},
				},
			},
		},
		{
			"MediansOfMedians",
			MediansOfMedians.Select,
			[]struct {
				index int
				initial, expected []int
			} {
				{
					3,
					[]int{65, 28, 59, 33, 21},
					[]int{21, 28, 33, 59, 65},
				},
			},
		},
	}

	for _, tc := range(testCases) {
		for _, e := range(tc.examples) {
			sortedCopy := append([]int{}, e.initial...)
			sort.Sort(sorter(sortedCopy))
			tc.algorithm(sorter(e.initial), e.index, 0, len(e.initial) - 1)
			assert.Equal(t, e.expected, e.initial)
			assert.Equal(t, e.initial[e.index], sortedCopy[e.index])
		}
	}
}


func TestBucketsSize1(t *testing.T) {

	a := []int{65, 59, 33, 21, 56, 22, 95, 50, 12, 90, 53, 28, 77, 39}
	FloydRivest.Buckets(sorter(a), 1)
	expected := append([]int{}, a...)
	// If bucket size is 1 then it should be the same
	sort.Sort(sorter(expected))

	assert.True(t, reflect.DeepEqual(a, expected), "If buckets have size 1 we should be sorting")
}

func TestBucketsBig(t *testing.T) {
	size := 4000
	a := make([]int, size)
	for i := 0; i < size; i++ {
		a[i] = rand.Int()
	}
	sortedCopy := append([]int{}, a...)
	sort.Sort(sorter(sortedCopy))
	for i := 4; i < 300; i++ {
		shuffle(a)
		bucketSize := i
		FloydRivest.Buckets(sorter(a), bucketSize)
		nBuckets := size / bucketSize
		maxs := make([]int, nBuckets)
		mins := make([]int, nBuckets)
		// Compute bounds of buckets
		for j := 0; j < nBuckets; j++ {
			min, max := minmaxIntSlice(a[bucketSize * j : bucketSize * (j + 1)])
			maxs[j] = max
			mins[j] = min
		}
		for j := 0; j < nBuckets -1; j++ {
			assert.True(t, maxs[j] < mins[j + 1], "All elements from one bucket should be smaller than elements from the next")
		}
	}
}


func TestSelectVariousIndices(t *testing.T) {
	arr := []int{65, 28, 59, 33, 21, 56, 22, 95, 50, 12, 90, 53, 28, 77, 39}
	sortedCopy := append([]int{}, arr...)
	sort.Sort(sorter(sortedCopy))
	for i := 0; i < len(arr); i++ {
		shuffle(arr)
		copyBefore := append([]int{}, arr...)
		FloydRivest.Select(sorter(arr), i, 0, len(arr) - 1)
		assert.Equal(t, arr[i], sortedCopy[i], "Failed with values: ", i, copyBefore, arr)
	}
}

func shuffle(a []int) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}


type sorter []int
func (s sorter) Len() int {
	return len(s)
}

func (s sorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sorter) Less (i, j int) bool {
	return s[i] < s[j]
}

func minmaxIntSlice (s []int) (min, max int) {
	if (len(s) == 0) {
		return 0, 0
	}
	min = s[0]
	max = s[0]
	for _, e := range s {
		if e < min {
			min = e
		}
		if e > min {
			max = e
		}
	}
	return min, max
}
