package lct

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/spf13/cast"
)

// DNode 双向链表节点
type DNode struct {
	Key, Val   int
	Next, Prev *DNode
}

// NewDoubleLinkedList 创建双向链表
func NewDoubleLinkedList(keys, vals []int) (*DNode, error) {
	if len(keys) != len(vals) {
		return nil, errors.New("keys and vals length not equal")
	}
	if len(keys) == 0 {
		return nil, errors.New("keys and vals length is 0")
	}
	head := &DNode{Key: keys[0], Val: vals[0]}
	cur := head
	for i := 1; i < len(keys); i++ {
		cur.Next = &DNode{Key: keys[i], Val: vals[i]}
		cur.Next.Prev = cur
		cur = cur.Next
	}
	return head, nil
}

func (d *DNode) Print() {
	var (
		keyBuf bytes.Buffer
		valBuf bytes.Buffer
	)
	keyBuf.Write([]byte(cast.ToString(d.Key)))
	valBuf.Write([]byte(cast.ToString(d.Val)))
	for d.Next != nil {
		keyBuf.Write([]byte(" -> "))
		keyBuf.Write([]byte(cast.ToString(d.Next.Key)))
		valBuf.Write([]byte(" -> "))
		valBuf.Write([]byte(cast.ToString(d.Next.Val)))
		d = d.Next
	}
	fmt.Println("Key: ", keyBuf.String())
	fmt.Println("Val: ", valBuf.String())
}

// ReversePrint 顺着该节点的prev，向前打印
func (d *DNode) ReversePrint() {
	var (
		keyBuf bytes.Buffer
		valBuf bytes.Buffer
	)
	keyBuf.Write([]byte(cast.ToString(d.Key)))
	valBuf.Write([]byte(cast.ToString(d.Val)))
	for d.Prev != nil {
		keyBuf.Write([]byte(" -> "))
		keyBuf.Write([]byte(cast.ToString(d.Prev.Key)))
		valBuf.Write([]byte(" -> "))
		valBuf.Write([]byte(cast.ToString(d.Prev.Val)))
		d = d.Prev
	}
	fmt.Println("Key: ", keyBuf.String())
	fmt.Println("Val: ", valBuf.String())
}
