package llrb

import (
	"fmt"
	"github.com/yutuer/mydatastruct/tree"
)

type rbNode struct {
	isRed  bool
	key    tree.Item
	value  tree.Item
	parent *rbNode
	left   *rbNode
	right  *rbNode
}

func (node *rbNode) Left() tree.BinaryNode {
	return node.left
}

func (node *rbNode) Right() tree.BinaryNode {
	return node.right
}

func (node *rbNode) String() string {
	return fmt.Sprintf("R:%v, key:%d, value:%d", node.isRed, node.key, node.value)
}

func NewRbNode(key, value tree.Item) *rbNode {
	return &rbNode{isRed: true, key: key, value: value}
}

func isRed(rbNode *rbNode) bool {
	return rbNode != nil && rbNode.isRed
}
