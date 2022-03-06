package linkedlist

// a: 1,1,1, 2,2,3
// b: 5,6,6,7
func removeSameValueNode(aList *LinkedList, bList *LinkedList) {
	// 保证2个有序列表都不为空
	if aList.len == 0 || bList.len == 0 {
		return
	}

	aNode := aList.head
	bNode := bList.head

	for aNode != nil {
		// b 结束了
		if bNode == nil {
			break
		}
		if aNode.val < bNode.val {
			aNode = aNode.next
			continue
		}
		if aNode.val > bNode.val {
			bNode = bNode.next
			continue
		}

		// a 节点和 b 节点 value 一样
		removeSameNodeAfterMe(aNode)
		removeSameNodeAfterMe(bNode)
		aNextNode := aNode.next
		bNextNode := bNode.next
		removeNode(aList, aNode)
		removeNode(bList, bNode)
		aNode = aNextNode
		bNode = bNextNode
	}
}

func removeSameNodeAfterMe(node *Node) {
	for node.next != nil {
		next := node.next
		if node.val != next.val {
			return
		}
		if next.next == nil {
			// 删除next即可
			node.next = nil
			return
		}
		next.next.before = node
		node.next = next.next
		removeSameNodeAfterMe(node)
	}
}
