package llrb

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/yutuer/mydatastruct/tree"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// 左倾斜红黑树, 不会有右边的红色
// 就只需要看最后插入的三个节点的关系就行了!!!! 因为234树 只会插入到叶子节点上面, 最多就是3个. 但是要注意, 可能会需要递归旋转
type llrbTree struct {
	root    *rbNode
	length  int
	compare func(k1, k2 tree.Item) int
}

func NewRbTree() tree.TreeOpera {
	return &llrbTree{compare: tree.CompareFunc}
}

// https://www.jianshu.com/p/0319d7781814
// 关键是看插入的是2node(1个值) 还是3node(2个值) , 对应5种不同的情况
func (tree *llrbTree) Add(k, v tree.Item) error {
	rbn := NewRbNode(k, v)
	if tree.length == 0 {
		rbn.isRed = false
		tree.root = rbn
	} else {
		for n := tree.root; n != nil; {
			r := tree.compare(k, n.key)

			if r != 0 {
				var node *rbNode
				if r < 0 {
					node = n.left
				} else {
					node = n.right
				}

				if node == nil {
					if r < 0 {
						n.left = rbn
					} else {
						n.right = rbn
					}

					rbn.parent = n
					circle(tree, n)
					break
				} else {
					n = node
				}
			} else {
				return errors.New(fmt.Sprintf("has key:%d", k))
			}

		}
	}

	tree.length++

	return nil
}

func (*llrbTree) Delete(key tree.Item) error {
	panic("implement me")
}

func (*llrbTree) Update(key, v tree.Item) error {
	panic("implement me")
}

func (tree *llrbTree) Get(key tree.Item) (tree.Item, error) {
	for n := tree.root; n != nil; {
		r := tree.compare(key, n.key)

		if r < 0 {
			n = n.left
		} else if r > 0 {
			n = n.right
		} else {
			return n.value, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("not find key:%d", key))
}

func (tree *llrbTree) String() string {
	panic("implement me")
}

func (tree *llrbTree) Header() tree.BinaryNode {
	return tree.root
}
