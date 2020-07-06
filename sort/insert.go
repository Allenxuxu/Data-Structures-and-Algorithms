package sort

func InsertSort(data Interface) {
	for i := 1; i < data.Len(); i++ {
		for j := i - 1; j >= 0; j-- {
			if data.Less(j, j+1) {
				break
			}

			data.Swap(j, j+1)
		}
	}
}
