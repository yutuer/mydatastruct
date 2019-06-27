package queue

type LimitIntQueue interface {
	IntQueue
	IsFull() bool
	Cap() int
}

type queue struct {
	data []int
}

type limitQueue struct {
	IntQueue
	capacity int
}

func NewLimitIntQueue(c int) LimitIntQueue {
	return &limitQueue{IntQueue: &queue{data: make([]int, 0)}, capacity: c}
}

func (lq *limitQueue) IsFull() bool {
	return lq.Size() == lq.capacity
}

func (lq *limitQueue) Cap() int {
	return lq.capacity
}
