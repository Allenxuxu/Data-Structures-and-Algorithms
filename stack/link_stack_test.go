package stack

import "testing"

func TestLinkStack(t *testing.T) {
	s := NewLinkStack()

	s.Push(1)
	if s.Length() != 1 {
		t.Fatal()
	}
	v, ok := s.Peek()
	if ok != true {
		t.Fatal()
	}
	if v.(int) != 1 {
		t.Fatal()
	}
	if s.Length() != 1 {
		t.Fatal()
	}
	tv, ok := s.Pop()
	if ok != true {
		t.Fatal()
	}
	if tv.(int) != 1 {
		t.Fatal()
	}
	if s.Length() != 0 {
		t.Fatal()
	}
	if _, ok := s.Peek(); ok != false {
		t.Fatal()
	}
	if _, ok := s.Pop(); ok != false {
		t.Fatal()
	}

	for i := 0; i < 4; i++ {
		s.Push(i)
	}
	if s.Length() != 4 {
		t.Fatal()
	}
	for i := 3; i >= 0; i-- {
		v, ok := s.Peek()
		if ok != true {
			t.Fatal()
		}
		if v.(int) != i {
			t.Fatal()
		}

		tv, ok := s.Pop()
		if ok != true {
			t.Fatal()
		}
		if tv.(int) != i {
			t.Fatal()
		}
	}
	if s.Length() != 0 {
		t.Fatal()
	}

	if _, ok := s.Peek(); ok != false {
		t.Fatal()
	}
	if _, ok := s.Pop(); ok != false {
		t.Fatal()
	}
}
