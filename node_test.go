package lct

import "testing"

func TestNode_Print(t *testing.T) {
	head := NewLinkedList([]int{0, 1, 2})
	head.Print()
}

func TestNode_ReversePrint(t *testing.T) {
	head := NewLinkedList([]int{0, 1, 2})
	head.ReversePrint()
}
