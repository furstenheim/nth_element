package project_test


import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"github.com/furstenheim/nth_element/utils"
	"github.com/furstenheim/nth_element/FloydRivest"
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

