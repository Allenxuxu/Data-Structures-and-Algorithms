package hashtable

import (
	"strconv"
	"testing"
)

func TestHashTable(t *testing.T) {
	ht := New(10)
	if ht.Length() != 0 {
		t.Fatal()
	}

	for i := 0; i < 10; i++ {
		ht.Add(strconv.Itoa(i), i)
	}
	for i := 0; i < 10; i++ {
		v, ok := ht.Get(strconv.Itoa(i))
		if !ok {
			t.Fatal("get ", i)
		}

		if v.(int) != i {
			t.Fatal()
		}
	}

	for i := 0; i < 10; i++ {
		if !ht.Set(strconv.Itoa(i), i+1) {
			t.Fatal()
		}
	}
	for i := 0; i < 10; i++ {
		v, ok := ht.Get(strconv.Itoa(i))
		if !ok {
			t.Fatal("get ", i)
		}

		if v.(int) != i+1 {
			t.Fatal()
		}
	}

	if ht.Length() != 10 {
		t.Fatal()
	}

	for i := 0; i < 5; i++ {
		ht.Del(strconv.Itoa(i))
	}
	for i := 0; i < 5; i++ {
		_, ok := ht.Get(strconv.Itoa(i))
		if ok {
			t.Fatal("get ", i)
		}
	}
	for i := 5; i < 10; i++ {
		v, ok := ht.Get(strconv.Itoa(i))
		if !ok {
			t.Fatal("get ", i)
		}

		if v.(int) != i+1 {
			t.Fatal()
		}
	}

	if ht.Length() != 5 {
		t.Fatal()
	}
}

func TestHashTable_resize(t *testing.T) {
	ht := New(5)
	if len(ht.data) != 5 {
		t.Fatal()
	}

	for i := 0; i < 5; i++ {
		ht.Add(strconv.Itoa(i), i)
	}
	for i := 0; i < 5; i++ {
		v, ok := ht.Get(strconv.Itoa(i))
		if !ok {
			t.Fatal("get ", i)
		}

		if v.(int) != i {
			t.Fatal()
		}
	}
	if len(ht.data) != 5 {
		t.Fatal()
	}

	ht.Add("5", 5)
	if len(ht.data) != 10 {
		t.Fatal()
	}

	for i := 6; i < 11; i++ {
		ht.Add(strconv.Itoa(i), i)
	}
	for i := 0; i < 11; i++ {
		v, ok := ht.Get(strconv.Itoa(i))
		if !ok {
			t.Fatal("get ", i)
		}

		if v.(int) != i {
			t.Fatal()
		}
	}
	if len(ht.data) != 20 {
		t.Fatal()
	}
	if ht.Length() != 11 {
		t.Fatal()
	}
}
