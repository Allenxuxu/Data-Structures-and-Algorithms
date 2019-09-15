package queue

import "testing"

func TestLinkQueue(t *testing.T) {
	queue := NewLinkQueue()

	{
		_, ok := queue.Pop()
		if ok != false {
			t.Fatal()
		}
		queue.Push(1)
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
			queue.Push(i)
		}

		queue.Push(10)

		for i := 0; i < 10; i++ {
			v, ok := queue.Pop()
			if ok != true {
				t.Fatal()
			}
			if v != i {
				t.Fatal()
			}
		}

		v, ok := queue.Pop()
		if ok != true {
			t.Fatal()
		}
		if v != 10 {
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

		queue.Clear()
		if queue.Length() != 0 {
			t.Fatal()
		}
	}
}
