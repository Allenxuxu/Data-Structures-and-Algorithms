package sort

import "testing"

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestSortIntSlice(t *testing.T) {
	data := ints
	a := IntSlice(data[0:])
	HeapSort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestSelectSort(t *testing.T) {
	data := ints
	a := IntSlice(data[0:])
	SelectSort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestInsertSort(t *testing.T) {
	data := ints
	a := IntSlice(data[0:])
	InsertSort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestInsertSort1(t *testing.T) {
	data := ints
	a := IntSlice(data[0:])
	InsertSort1(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestBubbleSort(t *testing.T) {
	data := ints
	a := IntSlice(data[0:])
	BubbleSort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestBubbleSort1(t *testing.T) {
	data := ints
	a := IntSlice(data[0:])
	BubbleSort1(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

// IsSorted reports whether data is sorted.
func IsSorted(data Interface) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}
