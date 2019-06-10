package maps

import (
	"testing"
)

func TestMap(t *testing.T) {
	const count = 50000
	m := NewMap(4)
	for i := 1; i < count; i++ {
		m.Put(i, i)
	}

	for i := 1; i < count; i += 2 {
		m.Remove(i)
	}

	for i := 1; i < count; i += 2 {
		if m.ContainKey(i) {
			t.Error(i)
		}
	}
}

func TestPow2(t *testing.T) {
	j := ToPow2(1)
	if j != 8 {
		t.Error("j != 8")
	}

	j = ToPow2(16)
	if j != 16 {
		t.Error("j != 16")
	}

	j = ToPow2(9)
	if j != 16 {
		t.Error("j != 16")
	}
}
