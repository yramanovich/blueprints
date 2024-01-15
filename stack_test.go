package blueprints

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	nums := []int{1, 2, 3}
	numsStack := stack[int](nums)

	assert.False(t, numsStack.empty())
	assert.Equal(t, 3, numsStack.len())
	assert.Equal(t, 3, numsStack.peek())
	assert.NotPanics(t, func() { numsStack.push(4) })
	assert.Equal(t, 4, numsStack.pop())
	assert.Equal(t, 3, numsStack.len())
}
