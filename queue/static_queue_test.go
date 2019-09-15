package queue

import "testing"

func TestStaticQueue(t *testing.T) {
	queue := NewStaticQueue(10)

	{
		_, ok := queue.Pop()
		if ok != false {
			t.Fatal()
		}
		if ok := queue.Push(1); ok != true {
			t.Fatal()
		}
		if queue.Length() != 1 {
			t.Fatal()
		}
		v, ok := queue.Pop()
		if ok != true {
			t.Fatal()
		}
		if v != 1 {
			t.Fatal()
		}
	}

	{
		for i := 0; i < 10; i++ {
			if ok := queue.Push(i); ok != true {
				t.Fatal()
			}
		}

		if ok := queue.Push(10); ok != false {
			t.Fatal()
		}

		for i := 0; i < 10; i++ {
			v, ok := queue.Pop()
			if ok != true {
				t.Fatal()
			}
			if v != i {
				t.Fatal()
			}
		}

		_, ok := queue.Pop()
		if ok != false {
			t.Fatal()
		}
	}

	{
		queue.Push('a')
		queue.Push('b')
		queue.Push('c')
		if queue.Length() != 3 {
			t.Fatal()
		}
		if queue.Capacity() != 10 {
			t.Fatal()
		}

		queue.Clear()
		if queue.Length() != 0 {
			t.Fatal()
		}
		if queue.Capacity() != 10 {
			t.Fatal()
		}
	}
}
