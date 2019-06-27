package queue

import "github.com/pkg/errors"

type priorIntQueue struct {
	data []int
}

func (pq *priorIntQueue) Push(i int) error {
	pq.data = append(pq.data, i)
	pq.up(pq.Size() - 1)
	return nil
}

func (pq *priorIntQueue) Poll() (int, error) {
	if pq.IsEmpty() {
		return -1, errors.New("no element")
	}

	ret := pq.data[0]
	pq.data[0] = pq.data[pq.Size()-1]
	pq.data = pq.data[:pq.Size()-1]
	pq.sink(0)
	return ret, nil
}

func (pq *priorIntQueue) IsEmpty() bool {
	return pq.Size() == 0
}

func (pq *priorIntQueue) Size() int {
	return len(pq.data)
}

func (pq *priorIntQueue) sink(index int) {
	if index >= pq.Size() {
		return
	}

	for ; index < pq.Size(); {
		var min = 0

		left := index*2 + 1
		if left >= pq.Size() {
			break
		}

		right := index*2 + 2
		if right >= pq.Size() {
			min = left
		} else {
			if pq.data[left] < pq.data[right] {
				min = left
			} else {
				min = right
			}
		}

		if pq.data[index] > pq.data[min] {
			pq.swap(index, min)
			index = min
		} else {
			break
		}
	}
}

func (pq *priorIntQueue) up(index int) {
	if index <= 0 {
		return
	}

	half := (index - 1) / 2
	for ; index > 0 && pq.data[index] < pq.data[half]; {
		pq.swap(index, half)
		index = half
		half = (index - 1) / 2
	}
}

func (pq *priorIntQueue) swap(i1 int, i2 int) {
	tmp := pq.data[i1]
	pq.data[i1] = pq.data[i2]
	pq.data[i2] = tmp
}

func NewPriorIntQueue() IntQueue {
	return &priorIntQueue{data: make([]int, 0)}
}
