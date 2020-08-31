package sort

func QuickSort(a []int) {
	Partition(a, 0, len(a)-1)
}
func Partition(a []int, low, height int) {
	if low < height {
		stand := getStandard(a, low, height)
		Partition(a, low, stand-1)
		Partition(a, stand+1, height)
	}
}

func getStandard(a []int, i, j int) int {
	index := a[i]

	for i < j {
		for i < j && a[j] >= index {
			j--
		}
		if a[j] < index {
			a[i] = a[j]
		}
		for i < j && a[i] <= index {
			i++
		}
		if a[i] > index {
			a[j] = a[i]
		}
	}
	a[i] = index
	return i
}

//冒泡排序，时间复杂读O(n2), 最好的情况是已经排好序，n ;最坏的情况是
func BubbleSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

//选择排序,n2,不稳定
func SelectSort(a []int) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}

//插入排序
func InsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := 0; j <= i; j++ {
			if a[j] > a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
}

func MergeSort(a []int) []int {
	if len(a) == 1 {
		return a
	}
	mid := len(a) / 2
	left := MergeSort(a[:mid])
	right := MergeSort(a[mid:])
	return Merge(left, right)
}

func Merge(a, b []int) []int {
	l, r := 0, 0
	result := []int{}
	for l < len(a) && r < len(b) {
		if a[l] < b[r] {
			result = append(result, a[l])
			l++
		} else {
			result = append(result, b[r])
			r++
		}
	}
	if l < len(a) {
		result = append(result, a[l:]...)
	}
	if r < len(b) {
		result = append(result, b[r:]...)
	}
	return result
}

//希尔排序
func ShellSort(a []int) {
	l := len(a)

	gap := l / 2

	for gap > 0 {
		for i := 1; i < len(a); i++ {
			for j := 0; j <= i; j++ {
				if a[j] > a[i] {
					a[j], a[i] = a[i], a[j]
				}
			}
		}
		gap = gap / 2
	}
}
