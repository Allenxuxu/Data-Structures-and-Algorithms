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
	head, tail := 0, len(data)-1
	for head < tail {
		if data[head] < mid {
			data[head], data[head] = data[head], data[head]
			head++
		} else {
			data[head], data[tail] = data[tail], data[head]
			tail--
		}
	}

	data[head] = mid
	QuickSort1(data[:head])
	QuickSort1(data[head+1:])
}
