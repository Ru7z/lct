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

// GenSortedLinkedList 生成测试有序列
func GenSortedLinkedList() *Node {
	l := GetSortedArray()
	return NewLinkedList(l)
}

// GenRandomLinkedList 生成测试随机列
func GenRandomLinkedList() *Node {
	l := GenRandomArray()
	return NewLinkedList(l)
}

// DetectCycle 是否有环, 返回入口
func (n *Node) DetectCycle() *Node {
	slow, fast := n, n
	for fast != nil {
		if fast.Next == nil {
			return nil
		} // 无环
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast { // 相遇后，fast重置为head，速度减半
			fast = n
			for {
				if slow == fast {
					fmt.Println("有环, 入口: ", fast.Val)
					return fast
				} // 找到
				slow, fast = slow.Next, fast.Next // 走下一步
			}
		}
	}
	return nil
}

// Print prints the list node.
func (n *Node) Print() {
	if n == nil {
		fmt.Println("nil")
		return
	}
	var (
		buf bytes.Buffer
	)
	buf.Write([]byte(cast.ToString(n.Val)))
	if entrance := n.DetectCycle(); entrance == nil {
		for n.Next != nil {
			buf.Write([]byte(" → "))
			buf.Write([]byte(cast.ToString(n.Next.Val)))
			n = n.Next
		}
	} else {
		fmt.Println("有环, 入口: ", entrance.Val)
		var (
			init = true
			head = n
		)
		for n.Next != entrance || init {
			if n.Next == entrance {
				init = false
			}
			buf.Write([]byte(" → "))
			buf.Write([]byte(cast.ToString(n.Next.Val)))
			n = n.Next
		}
		buf.Write([]byte("\n"))
		for head != entrance {
			buf.Write([]byte("    "))
			if head.Val >= 10 {
				buf.Write([]byte(" "))
			}
			head = head.Next
		}
		buf.Write([]byte("↑   "))
		if head.Val >= 10 {
			buf.Write([]byte(" "))
		}
		head = head.Next
		for head.Next != entrance {
			buf.Write([]byte("←   "))
			if head.Val >= 10 {
				buf.Write([]byte(" "))
			}
			head = head.Next
		}

		buf.Write([]byte("↓   "))
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
