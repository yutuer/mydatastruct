package _234tree

type Data []int

type Line []*Node

type Node struct {
	datas  Data
	childs Line
	level  int
}

func (this *Node) childLen() int {
	return len(this.childs)
}
func (this *Node) dataLen() int {
	return len(this.datas)
}

func (this *Node) add(v int) {
	if this.dataLen() == 2 {

	} else {
		// 直接加入
		this.datas = append(this.datas, v)
	}
}

func NewNode(v int, level int) *Node {
	n := &Node{datas: make([]int, 0), childs: make([]*Node, 0), level: level}
	n.datas = append(n.datas, v)
	return n
}

// 二三四树
type TTFTree struct {
	root  *Node
	level int
}

func (this *TTFTree) Add(v int) {
	if this.root == nil {
		this.root = NewNode(v, this.level)
	} else {
		this.root.add(v)
	}
}

