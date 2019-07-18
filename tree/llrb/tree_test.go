package llrb

import (
	"github.com/yutuer/mydatastruct/tree"
	"testing"
)

func TestInsert(t *testing.T) {
	rbTree := NewRbTree()

	d := []tree.String{"A", "S", "E", "R", "C", "D", "I", "N", "B", "X"}
	for _, v := range d {
		rbTree.Add(v, v)
	}

	tt := rbTree.(*llrbTree)

	if tt.length != len(d) {
		t.Errorf("tt.lenght should is 1000, but is %d", tt.length)
	}
}

func TestInsertAndDelete(t *testing.T) {
	//rbt := New()
	//
	//m := 0
	//n := 1000
	//for m < n {
	//	rbt.Insert(Int(m))
	//	m++
	//}
	//if rbt.Len() != uint(n) {
	//	t.Errorf("tree.Len() = %d, expect %d", rbt.Len(), n)
	//}
	//
	//for m > 0 {
	//	rbt.Delete(Int(m))
	//	m--
	//}
	//if rbt.Len() != 1 {
	//	t.Errorf("tree.Len() = %d, expect %d", rbt.Len(), 1)
	//}
}
