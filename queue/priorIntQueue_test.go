package queue

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	pq := NewPriorIntQueue()
	_, err := pq.Poll()
	if err == nil {
		t.Error("err is empty")
	}
}

func TestA(t *testing.T) {
	pq := NewPriorIntQueue()

	pq.Push(10)
	pq.Push(2)
	pq.Push(1)
	pq.Push(100)
	pq.Push(105)
	pq.Push(80)

	result := []int{1, 2, 10, 80, 100, 105}

	size := pq.Size()
	for i := 0; i < size; i++ {
		v, _ := pq.Poll()
		if result[i] != v {
			t.Error("expect value:", result[i], ", but:", v)
		}
	}

}
