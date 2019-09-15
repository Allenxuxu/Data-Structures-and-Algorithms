package list

import "testing"

func TestLinkList(t *testing.T) {
	linkList := LinkList{}

	{
		if linkList.Length() != 0 {
			t.Fatal()
		}
		if _, ok := linkList.Get(0); ok != false {
			t.Fatal()
		}

		for i := 0; i < 10; i++ {
			linkList.PushBack(i)
		}
		if linkList.Length() != 10 {
			t.Fatal()
		}
		linkList.Clear()
		if linkList.Length() != 0 {
			t.Fatal()
		}
	}

	{
		for i := 0; i < 10; i++ {
			linkList.PushBack(i)
		}

		for i := 9; i >= 0; i-- {
			v, ok := linkList.Get(i)
			if ok != true {
				t.Fatal()
			}
			if v.(int) != i {
				t.Fatal()
			}
			v, ok = linkList.Find(i)
			if ok != true {
				t.Fatal()
			}
			if v != i {
				t.Fatal()
			}
			v, ok = linkList.PopBack()
			if ok != true {
				t.Fatal()
			}
			if v.(int) != i {
				t.Fatal(i, v)
			}
		}

		_, ok := linkList.PopBack()
		if ok != false {
			t.Fatal()
		}

		if linkList.Length() != 0 {
			t.Fatal()
		}
	}

	{
		linkList.PushFront(0)
		v, ok := linkList.Get(0)
		if ok != true {
			t.Fatal()
		}
		if v.(int) != 0 {
			t.Fatal()
		}

		linkList.PushFront(-1)
		linkList.PushFront(-2)
		linkList.PushFront(-3)
		v, ok = linkList.PopFront()
		if ok != true {
			t.Fatal()
		}
		if v.(int) != -3 {
			t.Fatal()
		}
		v, ok = linkList.PopBack()
		if ok != true {
			t.Fatal()
		}
		if v.(int) != 0 {
			t.Fatal(v)
		}
		v, ok = linkList.Get(1)
		if ok != true {
			t.Fatal()
		}
		if v.(int) != -1 {
			t.Fatal()
		}
		v, ok = linkList.PopFront()
		if ok != true {
			t.Fatal()
		}
		if v.(int) != -2 {
			t.Fatal(v)
		}
		v, ok = linkList.PopFront()
		if ok != true {
			t.Fatal()
		}
		if v.(int) != -1 {
			t.Fatal()
		}

		if linkList.Length() != 0 {
			t.Fatal()
		}
	}
}
