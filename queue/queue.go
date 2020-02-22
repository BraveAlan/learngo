package queue

// 通过定义别名扩充系统类型
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int)) // 当v不是int时，只有在运行时才会知道
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
