package primary

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSlice(t *testing.T) {
	// t.Log(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	//t.Log(twoSum([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 6))
	//t.Log(plusOne([]int{8, 9, 9, 9, 9}))
	//t.Log(moveZeroes([]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}))

	// t.Log(rotate([]int{-1, -100, 3, 99}, 2))

	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	t.Log(isValidSudoku(board))

}

type MyIntint int

func TestQueue(t *testing.T) {
	t.Log(solveNQueens(5))
}

func TestGrammer(t *testing.T) {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestMergeSort(t *testing.T) {
	nums1 := []int{1, 2, 4, 5, 6, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	nums2 := []int{2, 3, 4, 65, 76, 78, 89}
	merge2(nums1, 6, nums2, 7)
	t.Log(nums1)
}
