package btree

import (
	"fmt"
	"log"
)

// 使用 主动分裂

// 数据节点
type dataNode struct {
	data  int
	left  *node
	right *node
	home  *node //属于哪个大node
}

func newDataNode(home *node, v int) *dataNode {
	return &dataNode{data: v, home: home}
}

// 大节点
type node struct {
	datas  []*dataNode
	parent *dataNode
}

func (node *node) dataLen() int {
	return len(node.datas)
}

func (node *node) onlyAdd(v int) int {
	l := node.dataLen()
	if l == 1 {
		n := node.datas[0]
		// 最底层的节点
		dataNode := newDataNode(node, v)
		if v < n.data {
			//放前面
			node.datas = append(node.datas[0:0], node.datas[1:]...)
			node.datas[0] = dataNode
			return 0
		} else if v > n.data {
			// 放后面
			node.datas = append(node.datas, dataNode)
			return 1
		} else {
			log.Fatalln("111")
		}
	} else if l == 2 {
		first := node.datas[0]

		// 最底层
		dataNode := newDataNode(node, v)

		second := node.datas[1]
		if v < first.data {
			// 插入到最前面
			node.datas = append(node.datas[0:1], node.datas[:]...)
			node.datas[0] = dataNode
			return 0
		} else if v > first.data && v < second.data {
			node.datas = append(node.datas[0:1], node.datas[:]...)
			node.datas[1] = dataNode
			return 1
		} else if v > second.data {
			node.datas = append(node.datas, dataNode)
			return 2
		} else {
			log.Fatalln("222")
		}
	}

	log.Fatalln("333")
	return -1
}

func (node *node) add(v int) {
	l := node.dataLen()
	if l == 0 {
		dataNode := newDataNode(node, v)
		node.datas = append(node.datas, dataNode)
	} else if l == 1 {
		n := node.datas[0]
		if n.left == nil && n.right == nil {
			// 最底层的节点
			dataNode := newDataNode(node, v)
			if v < n.data {
				//放前面
				node.datas = append(node.datas[0:0], node.datas[1:]...)
				node.datas[0] = dataNode
			} else if v > n.data {
				// 放后面
				node.datas = append(node.datas, dataNode)
			} else {
				log.Println("111")
			}
		} else {
			if v < n.data {
				n.left.add(v)
			} else if v > n.data {
				n.right.add(v)
			} else {
				log.Println("111")
			}
		}
	} else if l == 2 {
		first := node.datas[0]
		if first.left == nil && first.right == nil {
			// 最底层
			dataNode := newDataNode(node, v)

			second := node.datas[1]
			if v < first.data {
				// 插入到最前面
				node.datas = append(node.datas[0:1], node.datas[:]...)
				node.datas[0] = dataNode
			} else if v > first.data && v < second.data {
				node.datas = append(node.datas[0:1], node.datas[:]...)
				node.datas[1] = dataNode
			} else if v > second.data {
				node.datas = append(node.datas, dataNode)
			}
		} else {
			// 递归
			second := node.datas[1]
			if v < first.data {
				first.left.add(v)
			} else if v > first.data && v < second.data {
				// first的右边和second的左边 引用应该相等
				first.right.add(v)
			} else if v > second.data {
				second.right.add(v)
			}
		}

		// 超出了长度3, 需要调整.
		// 将中间的值提到上一层节点
		// 上滤
		node.up()
	}
}

// 递归查找能插入值的节点
func (node *node) findNode(v int) *node {
	return nil
}

func (node *node) up() {
	if len(node.datas) != 3 {
		return
	}

	middle := node.datas[1]
	if middle.home.parent == nil {
		// 没有父节点

		// 1.将middle上提成parent节点
		// 2.左右都设置为新的node, 并放置到新的node中
		// 3. 设置左右节点

		// 1
		middleNode := NewNode(nil, middle.data)

		// 2
		left := NewNode(middleNode.datas[0], node.datas[0].data)
		right := NewNode(middleNode.datas[0], node.datas[2].data)

		// 3
		middleNode.datas[0].left = left
		middleNode.datas[0].right = right
	} else {
		index := middle.home.parent.home.onlyAdd(middle.data)

		parentMiddle := middle.home.parent.home.datas[index]
		fmt.Println(parentMiddle)

	}
}

func NewNode(parent *dataNode, v int) *node {
	n := &node{datas: make([]*dataNode, 0), parent: parent}
	n.add(v)
	return n
}

// 二三四树
type TTFTree struct {
	root  *node
	level int
}

//Add 添加一个值
func (tree *TTFTree) Add(v int) {
	if tree.root == nil {
		tree.root = NewNode(nil, v)
	} else {
		n := tree.root.findNode(v)
		n.add(v)
	}
}

//Remove 移除一个值
func (tree *TTFTree) Remove(v int) bool {
	return false
}

//Contains 是否包含某个值
func (tree *TTFTree) Contains(v int) bool {
	return false
}

type NoChildNode interface {
	Add(v int)
}
type OneChildNode interface {
	Add(v int)
}
type TwoChildNode interface {
	Add(v int)
}
