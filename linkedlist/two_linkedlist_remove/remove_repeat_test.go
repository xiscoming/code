package linkedlist

import (
	"testing"
)

func Test_makeLinkedList(t *testing.T) {
	type args struct {
		linkedList *LinkedList
		data       []int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "a",
			args: args{
				linkedList: &LinkedList{},
				data:       []int{1, 2, 2, 3}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			makeLinkedList(tt.args.linkedList, tt.args.data)
			echoLinkedList("LinkedList", tt.args.linkedList)
		})
	}
}

func TestRemoveRepeatNode(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveRepeatNode()
		})
	}
}
