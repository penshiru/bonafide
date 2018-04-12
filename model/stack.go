package model

import "errors"

type Stack struct {
	Data []interface{}
}

var ErrEmptyStack = errors.New("Stack.go : Stack is empty")

//NewStack returns new stack with size number
func NewStack(number uint) *Stack {
	return &Stack{Data: make([]interface{}, 0, number)}
}

//Len return the number of items in Stack
func (s *Stack) Len() int {
	return len(s.Data)
}

//Push adds a new elem at the top of the stack
func (s *Stack) Push(value interface{}) {
	s.Data = append(s.Data, value)
}

//Pop take the the top item out, if Stack is empty, will return ErrEmptyStack decleared above
func (s *Stack) Pop() (interface{}, error) {
	if s.Len() > 0 {
		rect := s.Data[s.Len()-1]
		s.Data = s.Data[:s.Len()-1]
		return rect, nil
	}
	return nil, ErrEmptyStack
}

//Peek checks the top item. Notice, this is like a pointer:
//tmp, _ := s.Peek(); tmp = 123;
//s.Pop() should return 123, nil.
func (s *Stack) Peek() (interface{}, error) {
	if s.Len() > 0 {
		return s.Data[s.Len()-1], nil
	}
	return nil, ErrEmptyStack
}
