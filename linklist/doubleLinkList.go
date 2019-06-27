package linklist

import (
	"bytes"
	"fmt"
)

type intDoubleNode struct {
	val  int
	next *intDoubleNode
	prev *intDoubleNode
}

func NewDoubleNode(v int) *intDoubleNode {
	return &intDoubleNode{val: v}
}

type dummyHeadNode struct {
	next *intDoubleNode
}
type dummyTailNode struct {
	prev *intDoubleNode
}

type intDoubleLinkList struct {
	head *dummyHeadNode
	tail *dummyTailNode
}

func NewDoubleLinkList() IntLinkList {
	return &intDoubleLinkList{&dummyHeadNode{}, &dummyTailNode{}}
}

func (dl *intDoubleLinkList) String() string {
	var str bytes.Buffer

	for n := dl.head.next; n != nil; n = n.next {
		str.WriteString(fmt.Sprintf("%d", n.val))
		str.WriteString(",")
	}
	return str.String()
}

func (dl *intDoubleLinkList) append(v int) IntLinkList {
	if dl != nil {
		if dl.head.next == nil && dl.tail.prev == nil {
			node := NewDoubleNode(v)
			dl.head.next = node
			dl.tail.prev = node
		} else {
			prev := dl.tail.prev
			node := NewDoubleNode(v)
			prev.next = node
			node.prev = prev

			dl.tail.prev = node
		}
	}
	return dl
}

func (dl *intDoubleLinkList) contains(v int) bool {
	for n := dl.head.next; n != nil; n = n.next {
		if n.val == v {
			return true
		}
	}
	return false
}

func (dl *intDoubleLinkList) remove(v int) IntLinkList {
	for n := dl.head.next; n != nil; n = n.next {
		if n.val == v {
			if n.prev == nil { // 头
				next := n.next
				next.prev = nil
				dl.head.next = next
			} else if n.next == nil { //尾
				dl.tail.prev = n.prev
				n.prev.next = nil
			} else {
				prev := n.prev
				next := n.next
				prev.next = next
				next.prev = prev
			}
		}
	}
	return dl
}
