package queue

import "errors"

type IntQueue interface {
	Push(i int) error
	Poll() (int, error)
	IsEmpty() bool
	Size() int
}

func NewIntQueue() IntQueue {
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