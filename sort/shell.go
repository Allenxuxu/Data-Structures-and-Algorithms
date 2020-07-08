package sort

func ShellSort1(data []int) {
	for d := len(data) / 2; d > 0; d /= 2 {
		//开始插入排序
		for i := d; i < len(data); i += d {
			//满足条件则插入
			for j := i - d; j >= 0; j -= d {
				if data[j] < data[j+d] {
					break
				}

				data[j], data[j+d] = data[j+d], data[j]
			}
		}
	}
}

func ShellSort(data []int) {
	d := len(data)
	for d > 1 {
		d = d/3 + 1

		for i := d; i < len(data); i++ {
			for j := i - d; j >= 0; j -= d {
				if data[j] < data[j+d] {
					break
				}

				data[j], data[j+d] = data[j+d], data[j]
			}
		}
	}
}
