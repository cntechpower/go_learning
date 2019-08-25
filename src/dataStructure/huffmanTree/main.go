package main

import (
	"container/heap"
	"fmt"
)

//MinHeap Start
type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	items := x.(*Item)
	items.index = n
	*pq = append(*pq, items)
}

func (pq *PriorityQueue) Pop() (x interface{}) {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)

}

//MinHeapEnd

//huffManTree Start
type huffManTree struct {
	left   *huffManTree
	right  *huffManTree
	parent *huffManTree
	value  int
	weight int
}

func NewHuffManTree(value, weight int) *huffManTree {
	return &huffManTree{
		left:   nil,
		right:  nil,
		parent: nil,
		value:  value,
		weight: weight,
	}
}

func (t *huffManTree) insertTree(value, weight int) *huffManTree {
	return nil
}

func (t *huffManTree) getWpl() int {
	return 0
}

func main() {
	items := map[string]int{
		"banana": 3,
		"appale": 2,
		"pear":   4,
	}
	i := 0
	pq := make(PriorityQueue, len(items))
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	item2 := &Item{
		value:    "Orange",
		priority: 1,
	}
	heap.Push(&pq, item2)
	pq.update(item2, item2.value, 5)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s \n", item.priority, item.value)
	}

}
