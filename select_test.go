package project_test


import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"github.com/furstenheim/nth_element/utils"
	"github.com/furstenheim/nth_element/FloydRivest"
 	"github.com/furstenheim/nth_element/QuickSelect"
	"github.com/furstenheim/nth_element/MediansOfMedians"
	"sort"
)

func TestSelect_BigArray (t * testing.T) {
	var a = make([]int, 20000)
	for i, _ := range(a) {
		a[i] = i / 3
	}
	nTests := 20
	for j:=0; j < nTests; j++ {
		rand.Shuffle(len(a), func (i, j int) { a[i], a[j] = a[j], a[i] })
		k := rand.Intn(len(a))
		FloydRivest.Select(nthElementUtils.IntSorter(a), k, 0, len(a) - 1)
		assert.Equal(t, a[k], k / 3)
		for i := 0; i < k; i++ {
			assert.LessOrEqual(t, a[i], k/3)
		}
		for i := k+1; i < len(a); i++ {
			assert.LessOrEqual(t, k / 3, a[i])
		}
	}
}

func TestSelect_RandomArray (t * testing.T) {
	var a = make([]int, 200000)
	for i, _ := range(a) {
		a[i] = rand.Int()
	}
	testCases := []struct {
		name string
		algorithm func (array sort.Interface, k, left, right int)
	} {
		{

			"FloydRivest",
			FloydRivest.Select,
		},
		{
			"QuickSelect",
			QuickSelect.Select,
		},
		{
			"MediansOfMedians",
			MediansOfMedians.Select,
		},
	}
	for _, tc := range(testCases) {
		nTests := 100
		for j := 0; j < nTests; j++ {
			rand.Shuffle(len(a), func (i, j int) { a[i], a[j] = a[j], a[i] })
			k := rand.Intn(len(a))
			tc.algorithm(nthElementUtils.IntSorter(a), k, 0, len(a) - 1)
			v := a[k]
			for i := 0; i < k; i++ {
				assert.LessOrEqual(t, a[i], v)
			}
			for i := k + 1; i < len(a); i++ {
				assert.LessOrEqual(t, v, a[i])
			}
		}
	}
}

