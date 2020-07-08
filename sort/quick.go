package sort

func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}

	p := 0 // 采用动态随机值会更好
	i, j := 0, len(data)-1
	for i < j {
		for i < j && data[j] > data[p] {
			j--
		}
		for i < j && data[i] <= data[p] {
			i++
		}
		// swap
		data[i], data[j] = data[j], data[i]
	}

	// 此时 i == j
	data[i], data[p] = data[p], data[i]

	QuickSort(data[:i])
	QuickSort(data[i+1:])
}

func QuickSort1(data []int) {
	if len(data) <= 1 {
		return
	}

	mid := data[0]
	i, j := 0, len(data)-1
	for i < j {
		if data[i] < mid {
			i++
		} else {
			data[i], data[j] = data[j], data[i]
			j--
		}
	}

	data[i] = mid
	QuickSort1(data[:i])
	QuickSort1(data[i+1:])
}
