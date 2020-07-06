package sort

func BubbleSort(data Interface) {
	for i := 0; i < data.Len(); i++ {
		for j := data.Len() - 1; j > i; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}

// 优化：避免无效循环
func BubbleSort1(data Interface) {
	var isExchanged bool
	for i := 0; i < data.Len(); i++ {
		for j := data.Len() - 1; j > i; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
				isExchanged = true
			}
		}

		if !isExchanged { // 没有交换，说明已经是从小到大排序了，直接返回
			return
		}
	}
}
