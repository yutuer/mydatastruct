package linklist

import (
	"bytes"
	"fmt"
)

type dummyNode struct {
	next *intNode
}

type intNode struct {
	val  int
	next *intNode
}

func (l *intNode) Val() int {
	return l.val
}

func (l *intNode) Next() *intNode {
	return l.next
}

func newIntNode(v int) *intNode {
	return &intNode{val: v}
}

type intLinkList struct {
	head *dummyNode
}

func (il *intLinkList) append(v int) IntLinkList {
	next := newIntNode(v)
	if il.head.next == nil {
		il.head.next = next
	} else {
		n := il.head.next
		for ; n != nil; n = n.next {
			if n.next == nil {
				break
			}
		}
		n.next = next
	}
	return il
}

func (il *intLinkList) contains(v int) bool {
	for n := il.head.next; n != nil; n = n.next {
		if n.val == v {
			return true
		}
	}
	return false
}

func (il *intLinkList) remove(v int) IntLinkList {
	if il == nil {
		return il
	}

	if il.head.next == nil {
		return il
	}

	if il.head.next.val == v {
		il.head.next = il.head.next.next
	} else {
		n_prev := il.head.next
		for n := il.head.next; n != nil; n = n.next {
			if n.val == v {
				n_prev.next = n.next
				break
			}
			n_prev = n
		}
	}
	return il
}

func (il *intLinkList) String() string {
	var str bytes.Buffer
	if il != nil {
		for n := il.head.next; n != nil; n = n.next {
			str.WriteString(fmt.Sprintf("%d", n.val))
			str.WriteString(",")
		}
	}
	return str.String()
}

func NewIntList() IntLinkList {
	return &intLinkList{&dummyNode{}}
}
