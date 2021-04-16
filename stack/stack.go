package stack

import "sync"

type Stack struct {
	s []interface{}
	l sync.Mutex
}

func New() *Stack {
	return &Stack{s: make([]interface{}, 0)}
}

func (s *Stack) Push(a interface{}) {
	s.l.Lock()
	defer s.l.Unlock()

	s.s = append(s.s, a)
}

func (s *Stack) Pop() interface{} {
	s.l.Lock()
	defer s.l.Unlock()

	l := len(s.s)
	if l == 0 {
		return nil
	}
	a := s.s[l-1]
	s.s = s.s[:l-1]
	return a
}

func (s *Stack) Top() interface{} {
	s.l.Lock()
	defer s.l.Unlock()

	l := len(s.s)
	if l == 0 {
		return nil
	}
	return s.s[l-1]
}

func (s *Stack) Len() int {
	return len(s.s)
}
