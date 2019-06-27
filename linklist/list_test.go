package linklist

import (
	"fmt"
	"testing"
)

func TestLinkList(t *testing.T) {
	list := NewIntList()
	for i := 1; i < 10; i++ {
		list.append(i)
	}
	fmt.Println(list)

	if list.contains(50) {
		t.Error("don't has ", 50)
	}

	if !list.contains(5) {
		t.Error("should contains ", 5)
	}

	fmt.Println(list.remove(5))
}

func TestNewDoubleLinkList(t *testing.T) {
	list := NewDoubleLinkList()
	for i := 1; i < 10; i++ {
		list.append(i)
	}
	fmt.Println(list)

	if list.contains(50) {
		t.Error("don't has ", 50)
	}

	if !list.contains(5) {
		t.Error("should contains ", 5)
	}

	fmt.Println(list.remove(5))
}
