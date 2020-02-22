package main

import (
	"fmt"

	"golang.org/x/tools/container/intsets"
	"imooc.com/sb/learngo/tree"
)

// 通过组合的方式扩展原有的包
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func testSparse() {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1))
	fmt.Println(s.Has(100))
}

func main() {
	// var root treeNode
	root := tree.Node{Value: 100}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Print()
	root.SetValue(10)
	root.Print()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	//var pRoot2 *treeNode
	//pRoot2.SetValue(9)
	//pRoot2.Print()

	fmt.Println()
	root.Travel()
	fmt.Println()
	myroot := myTreeNode{&root}
	myroot.postOrder()

	testSparse()

	fmt.Println()

	var root2 tree.Node
	root2 = tree.Node{Value: 3}
	root2.Left = &tree.Node{}
	root2.Right = &tree.Node{5, nil, nil}
	root2.Right.Left = new(tree.Node)
	root2.Left.Right = tree.CreateNode(2)
	root2.Right.Left.SetValue(4)
	root2.Travel()

	nodeCount := 0
	root2.TravelseFunc(func(n *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node Count:", nodeCount)

	c := root2.TraverseWithChannle()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode)
}
