package blueprints

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList(t *testing.T) {
	ll := &linkedList[int]{}
	assert.NotPanics(t, func() { ll.add(10) })
	assert.NotPanics(t, func() { ll.add(9) })
	assert.Equal(t, 2, ll.len())
	assert.NotPanics(t, func() { ll.remove(10) })
	assert.NotPanics(t, func() { ll.remove(9) })
	assert.Equal(t, 0, ll.len())
}
