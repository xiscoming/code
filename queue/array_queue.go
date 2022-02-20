// 固定长度数组实现队列
// 使用取模来判断位置，使用变量来保存长度

package main

import (
	"fmt"
)

type MyQueue struct {
	front  int
	rear   int
	length int
	max    int
	data   []int
}

func NewQueue(max int) *MyQueue {
	return &MyQueue{
		front:  0,
		rear:   0,
		length: 0,
		max:    max,
		data:   make([]int, max),
	}
}

func (m *MyQueue) Push(item int) error {
	if m.length == m.max {
		return fmt.Errorf("queue is overlimit, max is %d", m.max)
	}

	m.data[m.rear] = item
	m.rear = (m.rear + 1) % m.max
	m.length++
	return nil
}

func (m *MyQueue) Pop() (int, error) {
	if m.length == 0 {
		return 0, fmt.Errorf("my queue is empty")
	}

	item := m.data[m.front]
	m.front = (m.front + 1) % m.max
	m.length--
	return item, nil
}

func main() {
	queue := NewQueue(5)
	for i := 1; i < 6; i++ {
		if err := queue.Push(i); err != nil {
			fmt.Printf("Push %d fail %s \n", i, err.Error())
			return
		}
	}

	if err := queue.Push(6); err == nil {
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
