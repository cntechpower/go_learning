package main

import (
	"container/list"
	"fmt"
	"sync"
)

var ErrNoOldestNodeExist = fmt.Errorf("no oldest node exist")

type LruList interface {
	Len() int
	Capacity() int
	SetSize(int) int
	Put(int, interface{}) interface{}
	Get(int) interface{}
	Remove(int) interface{}
	RemoveOldest() (interface{}, error)
}

// LruListSimple is a simple lru implementation using list.List and map
// it used map to quick search key.
// time complexity:
// Len: O(1)
// Capacity: O(1)
// SetSize: O(N) -> O(Len - newSize)
// Put: O(1)
// Get: O(1)
// Remove: O(1)
// RemoveOldest: O(1)
type LruListSimple struct {
	list       *list.List
	elementMap map[int]*list.Element
	capacity   int
	mu         sync.RWMutex
}

type LruListNode struct {
	key   int
	value interface{}
}

func NewLruList(capacity int) LruList {
	return &LruListSimple{
		list: list.New(),
		//let elementMap starting with small size.
		//so make()'size is omitted.
		elementMap: make(map[int]*list.Element),
		capacity:   capacity,
	}
}

// Len returns the current lru real length occupied.
func (l *LruListSimple) Len() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.Len()
}

// Capacity returns the max size of lru.
func (l *LruListSimple) Capacity() int {
	return l.capacity
}

//SetSize dynamic adjustment the max capacity of lru.
//It support auto shrink when new capacity is smaller than current lru capacity.
//It returns the shrink capacity.
func (l *LruListSimple) SetSize(newSize int) int {
	shrinkSize := 0
	for l.list.Len() > newSize {
		//lock will hold in RemoveOldest, so ignore here.
		_, _ = l.RemoveOldest()
		shrinkSize++
	}
	l.capacity = newSize
	return shrinkSize
}

//Put put v into lru front(newest node).
//It returns the value of put.
//It will overwrite exist k.
func (l *LruListSimple) Put(key int, v interface{}) interface{} {
	if l.list.Len() == l.capacity {
		// remove oldest node and ignore error.
		// because l.capacity always >=0 so RemoveOldest() will not throw an error.
		_, _ = l.RemoveOldest()
	}
	l.mu.RLock()
	node, exist := l.elementMap[key]
	l.mu.RUnlock()
	if exist && node != nil {
		l.Remove(key)
	}
	l.mu.Lock()
	element := l.list.PushFront(&LruListNode{
		key:   key,
		value: v,
	})
	l.elementMap[key] = element
	l.mu.Unlock()
	return element.Value
}

//Get returns the value corresponding k and new key and put v into lru back(newest node).
//It returns nil when key is not exist in lru.
func (l *LruListSimple) Get(key int) interface{} {
	l.mu.RLock()
	element, exist := l.elementMap[key]
	l.mu.RUnlock()
	if exist && element != nil {
		l.moveElementToBack(element)
		return element.Value.(*LruListNode).value
	} else {
		return nil
	}
}

//Remove returns the value corresponding k and remove it from lru.
//It returns nil when key is not exist in lru.
func (l *LruListSimple) Remove(key int) interface{} {
	l.mu.RLock()
	element, exist := l.elementMap[key]
	l.mu.RUnlock()
	if exist && element != nil {
		delete(l.elementMap, key)
		return l.list.Remove(element)
	} else {
		return nil
	}
}

// RemoveOldest remove the oldest node in lru.
// It returns th oldest value, and return ErrNoOldestNodeExist when list is empty.
func (l *LruListSimple) RemoveOldest() (interface{}, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	//The l.firstNode must not be nil.
	//because list.list.Remove(nil) will panic.
	if l.list.Front() == nil {
		return nil, ErrNoOldestNodeExist
	}
	node := l.list.Front().Value.(*LruListNode)
	delete(l.elementMap, node.key)
	l.list.Remove(l.list.Front())
	return node.value, nil
}

// moveElementToBack moves the corresponding element into lru back(newest node).
// just private use, should not calling in outside.
func (l *LruListSimple) moveElementToBack(element *list.Element) {
	l.mu.Lock()
	l.list.MoveToBack(element)
	l.mu.Unlock()
}
