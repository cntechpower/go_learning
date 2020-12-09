package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LRU(t *testing.T) {
	lru := NewLruList(5)

	//when index in small than capacity-1, expectLen=index+1
	//when index is bigger than capacity-1, expectLen=capacity
	expectLen := func(index int) int {
		if index >= 5 {
			return 5
		} else {
			return index + 1
		}
	}
	for i := 0; i <= 10; i++ {
		lru.Put(i, i)
		assert.Equal(t, i, lru.Get(i))
		assert.Equal(t, expectLen(i), lru.Len())
		assert.Equal(t, 5, lru.Capacity())
	}
	//lru should be: [6,7,8,9,10] now.

	assert.Equal(t, 2 /*shrink size should be 2*/, lru.SetSize(3))
	//after resize to 3, lru should be: [8,9,10] now.
	assert.Equal(t, nil, lru.Get(6))
	assert.Equal(t, nil, lru.Get(7))
	assert.Equal(t, 8, lru.Get(8))
	assert.Equal(t, 9, lru.Get(9))
	assert.Equal(t, 10, lru.Get(10))
	assert.Equal(t, 3, lru.Len())
	assert.Equal(t, 3, lru.Capacity())

	assert.Equal(t, 8, lru.Get(8))
	//after access 8, lru should be: [9,10,8] now.

	removed, err := lru.RemoveOldest()
	//after remove oldest node, lru should be: [10,8] now.
	assert.Equal(t, 9, removed)
	assert.Equal(t, nil, err)
	assert.Equal(t, 8, lru.Get(8))
	assert.Equal(t, nil, lru.Get(9))
	assert.Equal(t, 10, lru.Get(10))
	assert.Equal(t, lru.Len(), 2)
	assert.Equal(t, lru.Capacity(), 3)
	//after this, lru should be: [8,10] now.

	removed, err = lru.RemoveOldest()
	//after remove oldest node, lru should be: [10] now.
	assert.Equal(t, 8, removed)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, lru.Get(8))
	assert.Equal(t, nil, lru.Get(9))
	assert.Equal(t, 10, lru.Get(10))
	assert.Equal(t, lru.Len(), 1)
	assert.Equal(t, lru.Capacity(), 3)

	removed, err = lru.RemoveOldest()
	//after remove oldest node, lru should be: [] now.
	assert.Equal(t, 10, removed)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, lru.Get(8))
	assert.Equal(t, nil, lru.Get(9))
	assert.Equal(t, nil, lru.Get(10))
	assert.Equal(t, lru.Len(), 0)
	assert.Equal(t, lru.Capacity(), 3)

	//lru is empty now, RemoveOldest should return ErrNoOldestNodeExist
	removed, err = lru.RemoveOldest()
	assert.Equal(t, nil, removed)
	assert.Equal(t, ErrNoOldestNodeExist, err)
}
