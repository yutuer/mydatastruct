package queue

import "github.com/pkg/errors"

type IQueue interface {
	Push(i int) error
	Poll() (int, error)
	IsEmpty() bool
	Size() int
}

type ILimitQueue interface {
	IQueue
	IsFull() bool
	Cap() int
}

type Queue struct {
	data []int
}

func NewQueue() IQueue {
	return &Queue{data: make([]int, 0)}
}

func (q *Queue) Push(i int) error {
	q.data = append(q.data, i)
	return nil
}

func (q *Queue) Poll() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("isEmpty")
	}

	r := q.data[0]
	q.data = q.data[1:]
	return r, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Size() int {
	return len(q.data)
}

type LimitQueue struct {
	IQueue
	capacity int
}

func NewLimitQueue(c int) ILimitQueue {
	return &LimitQueue{IQueue: &Queue{data: make([]int, 0)}, capacity: c}
}

func (lq *LimitQueue) IsFull() bool {
	return lq.Size() == lq.capacity
}

func (lq *LimitQueue) Cap() int {
	return lq.capacity
}
