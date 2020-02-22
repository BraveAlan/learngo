package tree

import "fmt"

func (node *Node) Travel() {

	node.TravelseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

// 这样封装后，可以做除了打印之外更多的事
func (node *Node) TravelseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TravelseFunc(f)
	f(node)
	node.Right.TravelseFunc(f)
}

func (node *Node) TraverseWithChannle() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TravelseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
