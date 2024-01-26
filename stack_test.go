package blueprints

import (
	"fmt"
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

func TestMonotonousStack(t *testing.T) {
	nums := []int{3, 7, 8, 4, 5, 9, 1, 2, 11}
	monotonous := MonotonousStack(nums)
	fmt.Println(monotonous)
}
