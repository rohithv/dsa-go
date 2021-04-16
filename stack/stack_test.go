package stack

import "testing"

func TestEmptyStack(t *testing.T) {
	s := New()

	a := s.Top()
	if a != nil {
		t.Errorf("Expected nil but got %v", a)
	}

	s = New()
	a = s.Pop()
	if a != nil {
		t.Errorf("Expected nil but got %v", a)
	}
}

func TestStack(t *testing.T) {
	s := New()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Len() != 3 {
		t.Errorf("Expected length = 3, got %v", s.Len())
		return
	}
	if s.Top() != 3 {
		t.Errorf("Expected 3 but got %v", s.Top())
	}
	a := s.Pop()
	b := s.Pop()
	c := s.Pop()

	if a != 3 || b != 2 || c != 1 {
		t.Errorf("Pop is in incorrrect order %v %v %v", a, b, c)
		return
	}

	if s.Len() != 0 {
		t.Error("Stack length has to be 0")
	}
	if s.Top() != nil || s.Pop() != nil {
		t.Error("Stack has to be empty")
	}

}

//TODO: Add testcase for concurrent stack updates by multiple go routines
