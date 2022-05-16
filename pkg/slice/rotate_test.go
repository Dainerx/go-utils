package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	rotate(x, 2)
	assert.Equal(t, []int{4, 5, 1, 2, 3}, x)

	x = []int{1, 2, 3, 4, 5}
	rotate(x, -1)
	assert.Equal(t, []int{2, 3, 4, 5, 1}, x)

	x = []int{1, 2, 3, 4, 5}
	rotate(x, -40)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, x)
}
