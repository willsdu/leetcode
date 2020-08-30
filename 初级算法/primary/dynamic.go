package primary

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

f(n)=f(n-1)+f(n-2)
*/
func climbStairs(n int) int {
	a, b := 0, 1
	for i := 1; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素）
// 返回其最大和。
func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	sum := nums[0]
	max := nums[0]
	for i := 1; i < l; i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if sum < 0 && nums[i] > 0 {
			sum = 0
		}
		sum += nums[i]
		if sum > max {
			max = sum
		}
	}
	return max
}

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

注意：你不能在买入股票前卖出股票。
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	maxP := 0
	p := 0
	for i := 1; i < len(prices); i++ {
		p = prices[i] - min
		if p > maxP {
			maxP = p
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return maxP
}

/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
*/
func rob(nums []int) int {
	return 0
}
