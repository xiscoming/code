// 使用双向链表模拟固定长度队列
//

package main

import "fmt"

type NodeQueue struct {
	max int
	len int
	top *Node
	end *Node
}

type Node struct {
	val  int
	pre  *Node
	next *Node
}

func NewNodeQueue(max int) *NodeQueue {
	a := &NodeQueue{
		max: max,
		top: nil,
		end: nil,
	}
	return a

}

func (n *NodeQueue) Push(item int) error {
	if n.len == n.max {
		return fmt.Errorf("push fail, overlimit")
	}
	node := &Node{
		val: item,
	}

	if n.len == 0 {
		n.top = node
	} else {
		node.pre = n.end
		n.end.next = node
	}
	n.end = node
	n.len++
	return nil
}

func (n *NodeQueue) Pop() (item int, err error) {
	if n.len == 0 {
		return 0, fmt.Errorf("nodeQueue empty")
	}

	item = n.top.val
	if n.top.next == nil {
		n.top = nil
		n.end = nil
	} else {
		newTop := n.top.next
		newTop.pre = nil
		n.top = newTop
	}

	n.len--
	return
}

func main() {
	queue := NewNodeQueue(5)
	for i := 1; i < 6; i++ {
		if err := queue.Push(i); err != nil {
			fmt.Printf("Push %d fail %s \n", i, err.Error())
			return
		}
	}

	if err := queue.Push(6); err != nil {
		fmt.Println("after Push 6 fail", err.Error())
	}

	fmt.Printf("myQueue %+v \n", queue)

	for i := 1; i < 6; i++ {
		item, err := queue.Pop()
		fmt.Printf("pop item %d, err %s \n", item, err)
	}

	item, err := queue.Pop()
	fmt.Printf("after pop item %d, err %s \n", item, err)
}
