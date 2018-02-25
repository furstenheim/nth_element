package FloydRivest

import (
	"testing"
	"sort"
	"reflect"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"fmt"
)

func TestBuckets(t *testing.T) {
	a := []int{65, 59, 33, 21, 56, 22, 95, 50, 12, 90, 53, 28, 77, 39}
	Buckets(sorter(a), 1)
	expected := append([]int{}, a...)
	// If bucket size is 1 then it should be the same
	sort.Sort(sorter(expected))

	assert.True(t, reflect.DeepEqual(a, expected), "")
	shuffle(a)

	Buckets(sorter(a), 1)
}

func TestSelectKnownArray(t *testing.T) {
	arr := []int{65, 28, 59, 33, 21, 56, 22, 95, 50, 12, 90, 53, 28, 77, 39}
	sortedCopy := append([]int{}, arr...)
	sort.Sort(sorter(sortedCopy))
	index := 8
	Select(sorter(arr), index, 0, len(arr) - 1)
	assert.True(t, reflect.DeepEqual(arr, []int{39, 28, 28, 33, 21, 12, 22, 50, 53, 56, 59, 65, 90, 77, 95}))
	assert.Equal(t, arr[index], sortedCopy[index])
}

func TestSelectKnownArray2(t *testing.T) {
	arr := []int{22, 33, 12, 95, 65, 28, 28, 77, 39, 21, 59, 50, 53, 56, 90}
	sortedCopy := append([]int{}, arr...)
	sort.Sort(sorter(sortedCopy))
	index := 3
	Select(sorter(arr), index, 0, len(arr) - 1)
	assert.True(t, reflect.DeepEqual(arr, []int{12, 21, 22, 28, 28, 33, 39, 50, 53, 56, 59, 65, 77, 90, 95}))
	assert.Equal(t, arr[index], sortedCopy[index])
}

func TestSelectVariousIndices(t *testing.T) {
	arr := []int{65, 28, 59, 33, 21, 56, 22, 95, 50, 12, 90, 53, 28, 77, 39}
	sortedCopy := append([]int{}, arr...)
	sort.Sort(sorter(sortedCopy))
	for i := 0; i < len(arr); i++ {
		shuffle(arr)
		copyBefore := append([]int{}, arr...)
		Select(sorter(arr), i, 0, len(arr) - 1)
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