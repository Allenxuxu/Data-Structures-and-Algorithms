package sort

func SelectSort(data Interface) {
	for i := 0; i < data.Len(); i++ {
		min := i
		for j := i + 1; j < data.Len(); j++ {
			if data.Less(j, min) {
				min = j
			}
		}

		if min != i {
			data.Swap(i, min)
		}
	}
}
