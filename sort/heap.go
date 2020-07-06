package sort

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

func down(data Interface, root, n int) {
	for {
		child := 2*root + 1 // left child
		if child >= n {
			break
		}
		if child+1 < n && data.Less(child, child+1) {
			// right = child+1
			child++
		}
		if data.Less(child, root) {
			return
		}
		data.Swap(root, child)
		root = child
	}
}

func HeapSort(data Interface) {
	n := data.Len()
	// Build heap with greatest element at top.
	for i := n/2 - 1; i >= 0; i-- {
		down(data, i, n)
	}

	// Pop elements, largest first, into end of data.
	for i := n - 1; i >= 0; i-- {
		data.Swap(0, i)
		down(data, 0, i)
	}
}
