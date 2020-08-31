package sort

import "testing"

func TestQuickTest(t *testing.T) {
	a := []int{34, 4, 5, 56, 6, 6, 67, 78, 8, 233, 44}
	// QuickSort(a)
	// BubbleSort(a)
	// SelectSort(a)
	// InsertSort(a)
	// a = MergeSort(a)
	ShellSort(a)
	t.Log(a)
}
