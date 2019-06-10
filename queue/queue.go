package queue

import "github.com/pkg/errors"

type Queue interface {
	Push(i int) error
	Poll() (int, error)
	IsEmpty() bool
	Size() int
}

type LimitQueue interface {
	Queue
	IsFull() bool
	Cap() int
}

type queue struct {
	data []int
}

func NewQueue() Queue {
	return &queue{data: make([]int, 0)}
}

func (q *queue) Push(i int) error {
	q.data = append(q.data, i)
	return nil
}

func (q *queue) Poll() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("isEmpty")
	}

	r := q.data[0]
	q.data = q.data[1:]
	return r, nil
}

func (q *queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *queue) Size() int {
	return len(q.data)
}

type limitQueue struct {
	Queue
	capacity int
}

func NewLimitQueue(c int) LimitQueue {
	return &limitQueue{Queue: &queue{data: make([]int, 0)}, capacity: c}
}

func (lq *limitQueue) IsFull() bool {
	return lq.Size() == lq.capacity
}

func (lq *limitQueue) Cap() int {
	return lq.capacity
}
