package primary

import (
	"strings"
)

/*
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

例1
给定 nums = [0,0,1,1,1,2,2,3,3,4],

函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

你不需要考虑数组中超出新长度后面的元素。

*/
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	j := 0
	for i := 1; i < l; i++ {
		if nums[i] > nums[j] {
			nums[j+1] = nums[i]
			j += 1
		}
	}
	return j + 1
}

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

*/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; i < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

/*
给定一个整数数组，判断是否存在重复元素。

如果任意一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
*/
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		_, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			return true
		}
	}
	return false
}

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

*/
func singleNumber(nums []int) int {
	n := 0
	for _, v := range nums {
		n = n ^ v
	}
	return n
}

//intersect 计算两个数组的交集
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	result := []int{}
	for _, v := range nums1 {
		n := m[v]
		m[v] = n + 1
	}
	for _, v := range nums2 {
		_, ok := m[v]
		if ok {
			result = append(result, v)
			n := m[v]
			if n--; n == 0 {
				delete(m, v)
			} else {
				m[v] = n
			}
		}
	}
	return result
}

/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。
*/
func plusOne(digits []int) []int {
	l := len(digits)
	n := 1
	for i := l - 1; i >= 0; i-- {
		digits[i] += n
		n = digits[i]
		if n >= 10 {
			digits[i] = 0
			n = 1
		} else {
			n = 0
			break
		}
	}
	if n > 0 {
		return append([]int{1}, digits...)
	}
	return digits
}

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
func moveZeroes(nums []int) []int {
	zero := -1
	nzero := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nzero = i
		} else {
			if zero < 0 {
				zero = i
			}
		}
		if nzero >= 0 && zero >= 0 && nzero > zero {
			nums[zero], nums[nzero] = nums[nzero], nums[zero]
			zero += 1
			nzero = -1
		}
	}
	return nums
}

//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
func rotate(nums []int, k int) []int {
	l := len(nums)
	if l <= 1 || k <= 0 || k%l == 0 {
		return nums
	}
	stageValue := nums[0]
	index := 0
	count := l
	start := 0
	for count > 0 {
		index = (index + k) % l
		nums[index], stageValue = stageValue, nums[index]
		count -= 1
		if index == start {
			index += 1
			start += 1
			stageValue = nums[start]
		}
	}
	return nums
}

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buy := prices[0]
	sumP := 0
	p := 0
	for i := 1; i < len(prices); i++ {
		p = prices[i] - buy
		if p > 0 {
			sumP += p
		}
		buy = prices[i]
	}
	return sumP
}

func isValidSudoku(board [][]byte) bool {
	//行
	for i := 0; i < 9; i++ {
		if !RepeatNumberArray(board[i]) {
			return false
		}
	}
	//列
	for i := 0; i < 9; i++ {
		rows := []byte{}
		for j := 0; j < 9; j++ {
			rows = append(rows, board[j][i])
		}
		if !RepeatNumberArray(rows) {
			return false
		}
	}
	//块
	for i := 0; i < 9; {
		for j := 0; j < 9; {
			blocks := []byte{}
			blocks = append(blocks, board[i][j], board[i][j+1], board[i][j+2])
			blocks = append(blocks, board[i+1][j], board[i+1][j+1], board[i+1][j+2])
			blocks = append(blocks, board[i+2][j], board[i+2][j+1], board[i+2][j+2])
			if !RepeatNumberArray(blocks) {
				return false
			}
			j += 3
		}
		i += 3
	}
	return true
}

//RepeatNumberArray   数组中的重复数字 false表示有重复
func RepeatNumberArray(datas []byte) bool {
	dMap := make(map[byte]int)
	count := 0
	for i := 0; i < len(datas); i++ {
		if datas[i] != '.' {
			count++
			dMap[datas[i]] = 1
		}
	}
	return len(dMap) == count
}

//设计一种算法，打印 N 皇后在 N × N 棋盘上的各种摆法，其中每个皇后都不同行、不同列，也不在对角线上。
//这里的“对角线”指的是所有的对角线，不只是平分整个棋盘的那两条对角线。
func solveNQueens(n int) [][]string {
	result := [][]string{}
	if n == 1 {
		return append(result, []string{"Q"})
	}
	if n < 4 {
		return result
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			qn := []int{0, i}
			one := solveOneQueens(n, 1, j, qn)
			result = append(result, one...)
		}

	}
	return result
}
func solveOneQueens(n, i, j int, qn []int) [][]string {
	result := [][]string{}

	if CheckValid(qn, i, j) {
		qn = append(qn, i, j)
		i += 1
		j = 0
	} else {
		return result
	}

	if n*2 == len(qn) {
		return append(result, GenQueens(qn))
	}
	for ; i < n; i++ {
		has := false
		for ; j < n; j++ {
			if CheckValid(qn, i, j) {
				has = true
				result = append(result, solveOneQueens(n, i, j, qn)...)
			}
		}
		if !has {
			return result
		}
		j = 0
	}
	return result
}

func CheckValid(qn []int, i, j int) bool {
	for k := 0; k < len(qn); k += 2 {
		row := qn[k] - i
		column := qn[k+1] - j
		//检查同行同列  //检查对角线
		if row == 0 || column == 0 || (row*row == column*column) {
			return false
		}
	}
	return true
}
func MatchValid(qn []int, i, j int) bool {
	for k := 0; k < len(qn); k += 2 {
		if qn[k] == i && qn[k+1] == j {
			return true
		}
	}
	return false
}
func GenQueens(qn []int) []string {
	l := len(qn)
	result := []string{}

	for i := 0; i < l/2; i++ {
		row := []string{}
		for j := 0; j < l/2; j++ {
			if MatchValid(qn, i, j) {
				row = append(row, "Q")
			} else {
				row = append(row, ".")
			}
		}
		rowStr := strings.Join(row, "")
		result = append(result, rowStr)
	}
	return result
}

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。



说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m] = nums2[i]
		m += 1
		InsertIntoSort(nums1, m)
	}
}

//InsertIntoSort  插入一个已经排好序的切片
func InsertIntoSort(nums []int, l int) {
	for i := l - 1; i > 0; i-- {
		if nums[i] < nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
}

//合并数组的第二种解法
func merge2(nums1 []int, m int, nums2 []int, n int) {
	total := n + m
	l := len(nums1)

	for i := l - 1; i >= 0; i-- {
		if m > 0 && n > 0 {
			if nums1[m-1] > nums2[n-1] {
				nums1[i] = nums1[m-1]
				m--
			} else if nums1[m-1] <= nums2[n-1] {
				nums1[i] = nums2[n-1]
				n--
			}
			continue
		}
		if n == 0 && m > 0 {
			nums1[i] = nums1[m-1]
			m--
		}
		if m == 0 && n > 0 {
			nums1[i] = nums2[n-1]
			n--
		}

		if m == 0 && n == 0 {
			break
		}
	}

	sub := l - total
	for i := 0; i < l; i++ {
		if i < total {
			nums1[i], nums1[i+sub] = nums1[i+sub], nums1[i]
		} else {
			nums1[i] = 0
		}
	}
}
