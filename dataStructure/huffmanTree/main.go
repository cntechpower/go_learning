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
type huffManNode struct {
	left     *huffManNode
	right    *huffManNode
	parent   *huffManNode
	value    string
	priority int
	index    int
}

func (h *huffManNode) getWpl() int {
	wpl := 0
	h.travel(func(n *huffManNode) error {
		if n.value != "" {
			old := n
			depth := 0
			for n.parent != nil {
				depth++
				n = n.parent
			}
			wpl = wpl + old.priority*depth
		}
		return nil
	})
	return wpl
}

func (h *huffManNode) travel(fn func(n *huffManNode) error) error {
	if h.left != nil {
		if err := h.left.travel(fn); err != nil {
			return err
		}
	}
	if err := fn(h); err != nil {
		return err
	}
	if h.right != nil {
		if err := h.right.travel(fn); err != nil {
			return err
		}
	}
	return nil

}

type PriorityHuffManNodeList []*huffManNode

func (pq PriorityHuffManNodeList) Len() int { return len(pq) }

func (pq PriorityHuffManNodeList) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

func (pq PriorityHuffManNodeList) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq *PriorityHuffManNodeList) Push(x interface{}) {
	n := len(*pq)
	items := x.(*huffManNode)
	items.index = n
	*pq = append(*pq, items)
}

func (pq *PriorityHuffManNodeList) Pop() (x interface{}) {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func main() {
	charList := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 50,
		"f": 30,
		"g": 80,
	}
	pq := make(PriorityHuffManNodeList, len(charList))
	pqIdx := 0
	for k, v := range charList {
		pq[pqIdx] = &huffManNode{
			value:    k,
			priority: v,
			index:    pqIdx,
		}
		pqIdx++
	}
	heap.Init(&pq)
	var c1, c2 *huffManNode
	for len(pq) > 1 {
		c1 = heap.Pop(&pq).(*huffManNode)
		c2 = heap.Pop(&pq).(*huffManNode)
		parent := &huffManNode{
			priority: c1.priority + c2.priority,
			left:     c1,
			right:    c2,
		}
		c1.parent = parent
		c2.parent = parent
		heap.Push(&pq, parent)
	}
	root := heap.Pop(&pq).(*huffManNode)
	fmt.Printf("root is %v\n", root)
	wpl := root.getWpl()
	fmt.Printf("WPL is %v\n", wpl)

}
