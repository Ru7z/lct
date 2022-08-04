package lct

import "testing"

func TestDNode_Print(t *testing.T) {
	head, _ := NewDoubleLinkedList([]int{0, 1, 2}, []int{0, 1, 2})
	head.Print()
}
