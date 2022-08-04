package lct

import (
	"bytes"
	"fmt"

	"github.com/spf13/cast"
)

// Node is a node in list node.
type Node struct {
	Val  int
	Next *Node
}

// NewLinkedList creates a new list node.
func NewLinkedList(vals []int) *Node {
	head := &Node{Val: vals[0]}
	cur := head
	for i := 1; i < len(vals); i++ {
		cur.Next = &Node{Val: vals[i]}
		cur = cur.Next
	}
	return head
}

// Print prints the list node.
func (n *Node) Print() {
	var (
		buf bytes.Buffer
	)

	buf.Write([]byte(cast.ToString(n.Val)))
	for n.Next != nil {
		buf.Write([]byte(" -> "))
		buf.Write([]byte(cast.ToString(n.Next.Val)))
		n = n.Next
	}
	fmt.Println(buf.String())
}

// ReversePrint prints the list node in reverse order.
func (n *Node) ReversePrint() {
	t := n.Reverse(n)
	t.Print()
}

// Reverse reverses the list node.
func (n *Node) Reverse(head *Node) *Node {
	if head == nil {
		return nil
	}

	cur := head
	for head.Next != nil {
		t := head.Next.Next
		head.Next.Next = cur
		cur = head.Next
		head.Next = t
	}
	return cur
}
