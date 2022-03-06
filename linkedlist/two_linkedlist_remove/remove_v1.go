package linkedlist

// 循环了2遍
func removeRepeatNode(a *LinkedList, b *LinkedList) {
	aNodeMap := map[int][]*Node{}
	node := a.head

	for node != nil {
		nodeList, ok := aNodeMap[node.val]
		if !ok {
			nodeList = []*Node{}
		}
		aNodeMap[node.val] = append(nodeList, node)

		if node.next == nil {
			break
		}
		node = node.next
	}

	// fmt.Println("for range a", aNodeMap)

	node = b.head

	for node != nil {
		nodeList, ok := aNodeMap[node.val]
		if !ok {
			if node.next == nil {
				break
			}
			node = node.next
			continue
		}

		for _, n := range nodeList {
			removeNode(a, n)
		}

		removeNode(b, node)

		if node.next == nil {
			break
		}
		node = node.next
	}

}
