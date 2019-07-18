package llrb

// 右边是红色结点, 左旋
func leftRotate(tree *llrbTree, n *rbNode) *rbNode {
	right := n.right

	left := right.left
	n.right = left
	if left != nil {
		left.parent = n
	}

	parent := n.parent
	right.parent = parent

	right.left = n
	n.parent = right

	if parent == nil {
		tree.root = right
	} else {
		if tree.compare(n.key, parent.key) < 0 {
			parent.left = right
		} else {
			parent.right = right
		}
	}

	// 更改颜色
	right.isRed = n.isRed
	n.isRed = true

	return right
}

// 右旋转上半部分
func rightRotate(tree *llrbTree, n *rbNode) *rbNode {
	parent := n.parent
	grandParent := parent.parent

	right := n.right
	parent.left = right
	if right != nil {
		right.parent = parent
	}

	n.right = parent
	parent.parent = n

	n.parent = grandParent

	n.isRed = parent.isRed
	parent.isRed = true

	if grandParent == nil {
		tree.root = n
	} else {
		if tree.compare(n.key, grandParent.key) < 0 {
			grandParent.left = n
		} else {
			grandParent.right = n
		}
	}
	return n
}
