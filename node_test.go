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

func TestGenRandomLinkedList(t *testing.T) {
	GenRandomLinkedList().Print()
}

func TestGenSortedLinkedList(t *testing.T) {
	GenSortedLinkedList().Print()
}
