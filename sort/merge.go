package sort

func MergeSort(data []int) {
	if len(data) <= 1 {
		return
	}

	tmp := make([]int, len(data))
	copy(tmp, data)
	a := tmp[:len(tmp)/2]
	b := tmp[len(tmp)/2:]
	MergeSort(a)
	MergeSort(b)

	merge(data, a, b)
}

func merge(array []int, a []int, b []int) {
	index, i, j := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			array[index] = a[i]
			i++
		} else {
			array[index] = b[j]
			j++
		}

		index++
	}

	for i < len(a) {
		array[index] = a[i]
		i++
		index++
	}
	for j < len(b) {
		array[index] = b[j]
		j++
		index++
	}
}
