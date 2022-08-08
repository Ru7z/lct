# lct

Go语言的一个LeetCode工具。

## 链表

```go
package service

import (
	"testing"

	l "github.com/ru7z/lct"
)

func Test_DetectCycle(t *testing.T) {
	h1 := l.GenRandomLinkedList()
	h := h1
	for h.Next != nil {
		h = h.Next
	}
	h.Next = h1.Next
	t.Log(h1.DetectCycle())
	h1.Print()
} 
```

```text
有环, 入口:  1
8 → 1 → 3 → 9 → 2 → 10
    ↑   ←   ←   ←   ↓
```