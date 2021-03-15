package main

//Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := []*Node{root}
	level := 0
	for len(queue) != 0 {
		checkpoint := len(queue)
		res = append(res, []int{})
		for i := 0; i < checkpoint; i++ {
			if queue[i] != nil {
				res[level] = append(res[level], queue[i].Val)
				for _, n := range queue[i].Children {
					queue = append(queue, n)
				}
			}
		}
		queue = queue[checkpoint:]
		level++
	}
	return res
}
