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

// 优化：减少交换
func InsertSort1(data []int) {
	for i := 1; i < len(data); i++ {
		tmp := data[i]
		index := i
		for j := i; j > 0; j-- {
			if data[j-1] < tmp {
				break
			}

			data[j] = data[j-1]
			index = j - 1
		}

		data[index] = tmp
	}
}
