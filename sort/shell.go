package sort

func ShellSort(data []int) {
	for step := len(data) / 2; step > 0; step /= 2 {
		//开始插入排序
		for i := step; i < len(data); i += step {
			//满足条件则插入
			for j := i - step; j >= 0; j -= step {
				if data[j] < data[j+step] {
					break
				}

				data[j], data[j+step] = data[j+step], data[j]
			}
		}
	}
}
