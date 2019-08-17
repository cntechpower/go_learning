package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

func (n *node) String() string {
	if n.left == nil && n.right == nil {
		return fmt.Sprintf("value is %v,left value is nil,right value is nil\n", n.value)
	} else if n.left == nil && n.right != nil {
		return fmt.Sprintf("value is %v,left value is nil,right value is %v\n", n.value, n.right)
	} else if n.left != nil && n.right == nil {
		return fmt.Sprintf("value is %v,left value is %v,right value is nil\n", n.value, n.left)
	}
	return fmt.Sprintf("\nvalue is %v,left value is %v,right value is %v\n", n.value, n.left, n.right)
}

func (n *node) PerOrderTraversal() {
	fmt.Println(n.value)
	if n.left != nil {
		n.left.PerOrderTraversal()
	}
	if n.right != nil {
		n.right.PerOrderTraversal()
	}

}

func (n *node) InOrderTraversal() {
	if n.left != nil {
		n.left.InOrderTraversal()
	}
	fmt.Println(n.value)
	if n.right != nil {
		n.right.InOrderTraversal()
	}
}

func (n *node) PostOrderTraversal() {
	if n.left != nil {
		n.left.PostOrderTraversal()
	}
	if n.right != nil {
		n.right.PostOrderTraversal()
	}
	fmt.Println(n.value)
}

func (n *node) LevelOrderTraversal() {
	buffer := make([]*node, 0)
	buffer = append(buffer, n)
	for len(buffer) != 0 {
		fmt.Println(buffer[0].value)
		if buffer[0].left != nil {
			buffer = append(buffer, buffer[0].left)
		}
		if buffer[0].right != nil {
			buffer = append(buffer, buffer[0].right)
		}
		buffer = buffer[1:]
	}

}

func (n *node) insert(v int) {
	newNode := &node{value: v, left: nil, right: nil}
	for true {
		if v == n.value {
			fmt.Printf("%v already exists\n", v)
			return
		} else if v > n.value && n.right != nil {
			n = n.right
		} else if v < n.value && n.left != nil {
			n = n.left
		} else if v > n.value && n.right == nil {
			n.right = newNode
			return
		} else if v < n.value && n.left == nil {
			n.left = newNode
			return
		}
	}

}

func newBinaryTree(v int) *node {
	return &node{value: v, left: nil, right: nil}
}

func main() {
	bt := newBinaryTree(20)
	bt.insert(10)
	bt.insert(30)
	bt.insert(5)
	bt.insert(15)
	bt.insert(25)
	bt.insert(35)
	bt.insert(2)
	bt.insert(7)
	bt.insert(12)
	bt.insert(17)
	bt.insert(22)
	bt.insert(27)
	bt.insert(33)
	bt.insert(38)
	fmt.Println("PerOrderTraversal")
	fmt.Println("-----------------------------------------")
	bt.PerOrderTraversal()
	fmt.Println("-----------------------------------------")
	fmt.Println("InOrderTraversal")
	fmt.Println("-----------------------------------------")
	bt.InOrderTraversal()
	fmt.Println("-----------------------------------------")
	fmt.Println("PostOrderTraversal")
	fmt.Println("-----------------------------------------")
	bt.PostOrderTraversal()
	fmt.Println("-----------------------------------------")
	fmt.Println("LevelOrderTraversal")
	fmt.Println("-----------------------------------------")
	bt.LevelOrderTraversal()
	fmt.Println("-----------------------------------------")

}
