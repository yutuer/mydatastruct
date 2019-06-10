package stack

import "errors"

type Stack interface {
	Size() int
	Push(i int)
	Pop() (int, error)
}

type stack struct {
	data []int
}

func NewStack() Stack {
	return &stack{data: make([]int, 0)}
}

func (s *stack) Size() int {
	return len(s.data)
}

func (s *stack) Push(i int) {
	s.data = append(s.data, i)
}

func (s *stack) Pop() (int, error) {
	if s.Size() == 0 {
		return -1, errors.New("stack is empty")
	}

	lastIndex := len(s.data) - 1
	r := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return r, nil
}
