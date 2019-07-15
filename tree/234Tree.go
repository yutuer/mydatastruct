package tree

type dataNode struct {
	data  int
	left  *node
	right *node
	home  *node //属于哪个大node
}

type node struct {
	datas  []*dataNode
	parent *dataNode
}

func (node *node) childLen() int {
	return len(node.childs)
}
func (node *node) dataLen() int {
	return len(node.datas)
}

func (node *node) add(v int) {
	if node.dataLen() == 0 {
		node.datas = append(node.datas, v)
	} else {
		// 找到应该插入的地方插入
		l := len(node.datas)
		for i := 0; i < len(node.datas); i++ {
			if v < node.datas[i] {
				l = i
				break
			}
		}

		d := node.datas[l:len(node.datas)]
		node.datas = append(node.datas[0:l], v)
		node.datas = append(node.datas, d...)

		// 超出了长度3, 需要调整.
		// 将中间的值提到上一层节点
		if node.dataLen() == 3 {

		}
	}
}

// 递归查找能插入值的节点
func (node *node) findNode(v int) *node {
	return nil
}

func NewNode(parent *node, v int) *node {
	n := &node{datas: make([]int, 0, 2), childs: make([]*node, 0, 3)}
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
