package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

// 自定义工厂函数
func CreateNode(value int) *Node {
	return &Node{Value: value} // 注意返回了局部变量的地址，不用考虑变量是创建在栈上还是堆上
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored.")
		return // 虽然nil指针能够安全传进来，但是执行下面语句的时候还是会报错
	}
	node.Value = value
}
