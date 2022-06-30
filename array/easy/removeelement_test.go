package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testData = []struct {
		data          []int
		elemToRemove  int
		expectedElems int
		expectedArray []int
	}{
		{
			data:          []int{3, 2, 2, 3},
			elemToRemove:  3,
			expectedElems: 2,
			expectedArray: []int{2, 2, 3, 3},
		},
		{
			data:          []int{0, 1, 2, 2, 3, 0, 4, 2},
			elemToRemove:  2,
			expectedElems: 5,
			expectedArray: []int{0, 1, 4, 0, 3, 2, 2, 2},
		},
	}
)

func TestRemoveElement(t *testing.T) {
	for _, v := range testData {
		actual := removeElement(v.data, v.elemToRemove)
		assert.Equal(t, v.expectedElems, actual)
		assert.ElementsMatch(t, v.data[:v.expectedElems], v.expectedArray[:v.expectedElems])
	}
}
