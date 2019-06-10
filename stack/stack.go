package stack

import "errors"

type IStack interface {
	Size() int
	Push(i int)
	Pop() (int, error)
}

type Stack struct {
	data []int
}

func NewStack() IStack {
	return &Stack{data: make([]int, 0)}
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Push(i int) {
	s.data = append(s.data, i)
}

func (s *Stack) Pop() (int, error) {
	if s.Size() == 0 {
		return -1, errors.New("stack is empty")
	}

	lastIndex := len(s.data) - 1
	r := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return r, nil
}
