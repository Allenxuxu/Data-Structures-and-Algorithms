package sort

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p IntSlice) HeapSort() { HeapSort(p) }

type Float64Slice []float64

func (p Float64Slice) Len() int           { return len(p) }
func (p Float64Slice) Less(i, j int) bool { return p[i] < p[j] || isNaN(p[i]) && !isNaN(p[j]) }
func (p Float64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// isNaN is a copy of math.IsNaN to avoid a dependency on the math package.
func isNaN(f float64) bool {
	return f != f
}

func (p Float64Slice) HeapSort() { HeapSort(p) }

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p StringSlice) HeapSort() { HeapSort(p) }
