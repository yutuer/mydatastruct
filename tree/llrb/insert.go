package llrb

func circle(tree *llrbTree, n *rbNode) {
	if isRed(n.right) && isRed(n.left) {
		// 因为子结点都是红色了, 再反转子结点的颜色
		revertTwoChildAllRed(tree, n)
	} else if isRed(n.right) {
		// 先左旋转
		n = leftRotate(tree, n)
	}

	if isRed(n) && isRed(n.left) {
		// 先右旋转上半部分
		n = rightRotate(tree, n)

		// 因为子结点都是红色了, 再反转子结点的颜色
		revertTwoChildAllRed(tree, n)
	}
}

// 同一个黑色父节点, 有2个红色的子结点(终结方法, 不需要返回下次需要处理的结点)
func revertTwoChildAllRed(tree *llrbTree, n *rbNode) {
	n.left.isRed = false
	n.right.isRed = false

	if n != tree.root {
		n.isRed = true

		parent := n.parent
		if tree.compare(n.key, parent.key) > 0 {
			// n 变成红色的了, 要根据parent的颜色来旋转
			circle(tree, parent)
		}
	}
}
