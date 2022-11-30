/*
 * @Author: TomaChen513
 * @Date: 2022-11-28 10:26:42
 * @LastEditors: TomaChen513
 * @LastEditTime: 2022-11-30 10:31:37
 * @FilePath: /LeetCode/CXYLeetCodeGo/DynamicProgramming/dynamicProgram.go
 * @Description:
 *
 * Copyright (c) 2022 by TomaChen513 xy_chen513@163.com, All Rights Reserved.
 */
package dynamicprogramming

// 188
// 动态规划关键在于设定状态
func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	status := make([]int, (2*k+1)*len(prices))

	for i := range dp {
		dp[i] = status[:2*k+1]
		status = status[2*k+1:]
	}
	for j := 1; j < 2*k; j += 2 {
		dp[0][j] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k; j += 2 {
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}
	return dp[len(prices)-1][2*k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 309
// 动态规划关键就是确定状态
func maxProfit4(prices []int) int {
	n := len(prices)
	if len(prices) < 2 {
		return 0
	}
	dp := make([][]int, len(prices))
	status := make([]int, len(prices)*4)

	for i := range dp {
		dp[i] = status[:4]
		status = status[4:]
	}
	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][1]-prices[i], dp[i-1][3]-prices[i]))
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}

	return max(dp[n-1][1], max(dp[n-1][2], dp[n-1][3]))
}

func maxProfit5(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
	}
	return dp[n-1][1]
}

// 300
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	maxNum:=0
	
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]>nums[j] {
				dp[i]=max(dp[i],dp[j]+1)
			}
		}
		maxNum=max(maxNum,dp[i])
	}
	return maxNum+1
}

// 674
func findLengthOfLCIS(nums []int) int {
	dp:=make([]int,len(nums))
	maxNum:=dp[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]>nums[i-1] {
			dp[i]=dp[i-1]+1
		}
		maxNum=max(maxNum,dp[i])
	}
	return maxNum+1
}
