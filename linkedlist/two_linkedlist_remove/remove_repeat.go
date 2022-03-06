package linkedlist

import "fmt"

// 两个有序链表去掉重复的元素

type Node struct {
	val    int
	next   *Node
	before *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func makeLinkedList(linkList *LinkedList, data []int) {
	max := len(data) - 1
	dataMap := map[int]*Node{}
	for i := max; i > -1; i-- {
		n := &Node{
			val:  data[i],
			next: nil,
		}
		dataMap[i] = n
		if i == max {
			continue
		}
		n.next = dataMap[i+1]
		n.next.before = n
	}

	linkList.head = dataMap[0]
	linkList.len = len(data)

	return
}

func echoLinkedList(prefix string, a *LinkedList) {
	list := []int{}
	if a == nil {
		fmt.Println(prefix, list)
		return
	}
	node := a.head
	for node != nil {
		list = append(list, node.val)
		if node.next == nil {
			break
		}
		node = node.next
	}
	fmt.Println(prefix, list)
}

// 删除元素
func removeNode(nodeList *LinkedList, n *Node) {
	if n == nil || nodeList == nil {
		return
	}

	// 唯一的元素
	if n.before == nil && n.next == nil {
		nodeList.head = nil
		return
	}

	// 首部的元素
	if n.before == nil {
		nodeList.head = n.next
		n.next.before = nil
		return
	}

	// 尾部
	if n.next == nil {
		n.before.next = nil
		return
	}

	n.before.next = n.next
	n.next.before = n.before

	return
}

func RemoveRepeatNode() {
	// a: 1,2,2,3
	a := &LinkedList{}
	makeLinkedList(a, []int{1, 1, 2, 2, 3, 3, 4, 4, 5})

	// b: 2,3,3,4
	b := &LinkedList{}
	makeLinkedList(b, []int{3, 3, 4, 4, 6})
	echoLinkedList("a linked list is ", a)
	echoLinkedList("b linked list is ", b)

	fmt.Println("---begin remove same node-----")

	// removeRepeatNode(a, b)
	removeSameValueNode(a, b)

	echoLinkedList("a linked list is ", a)
	echoLinkedList("b linked list is ", b)
}
