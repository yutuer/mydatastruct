package tree

import "fmt"

type TreeOpera interface {
	Add(key, v Item) error
	Delete(key Item) error
	Update(key, v Item) error
	Get(key Item) (Item, error)

	fmt.Stringer

	Header() BinaryNode
}

type BinaryNode interface {
	Left() BinaryNode
	Right() BinaryNode

	fmt.Stringer
}

type Item interface {
	Less(a Item) bool
}

func CompareFunc(a, b Item) int {
	if a.Less(b) {
		return -1
	} else if b.Less(a) {
		return 1
	} else {
		return 0
	}
}

type String string

func (s String) Less(a Item) bool {
	return s < a.(String)
}

type Byte byte

func (b Byte) Less(a Item) bool {
	return b < a.(Byte)
}

type Int int

func (i Int) Less(a Item) bool {
	return i < a.(Int)
}
