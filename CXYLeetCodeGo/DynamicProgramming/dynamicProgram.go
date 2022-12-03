/*
 * @Author: TomaChen513
 * @Date: 2022-11-28 10:26:42
 * @LastEditors: TomaChen513
 * @LastEditTime: 2022-11-30 15:19:41
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

// 718
func findLength(nums1 []int, nums2 []int) int {
	dp:=make([][]int,len(nums1))
	subDp:=make([]int,len(nums1)*len(nums2))
	for i := 0; i < len(nums1); i++ {
		dp[i] = subDp[:len(nums2)]
		subDp=subDp[len(nums2):]
	}
	maxNum:=0

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i]==nums2[j] {
				dp[i][j]=1
				maxNum=1
			}
		}
	}


	for i := 1; i < len(nums1); i++ {
		for j := 1; j < len(nums2); j++ {
			if nums1[i]==nums2[j] {
				dp[i][j]=dp[i-1][j-1]+1
				maxNum=max(maxNum,dp[i][j])
			}
		}
	}

	return maxNum
}

// 1143
// func longestCommonSubsequence(text1 string, text2 string) int {
// 	text1Bytes:=[]byte(text1)
// 	text2Bytes:=[]byte(text2)
// 	dp:=make([][]int,len(text1Bytes))
// 	subDp:=make([]int,len(text1Bytes)*len(text2Bytes))
// 	for i := 0; i < len(text1Bytes); i++ {
// 		dp[i] = subDp[:len(text2Bytes)]
// 		subDp=subDp[len(text1Bytes):]
// 	}
// 	// maxNum:=0

// 	for i := 0; i < len(text1Bytes); i++ {
// 		for j := 0; j < len(text2Bytes); j++ {
// 			if text1Bytes[i]==text2Bytes[j] {
// 				dp[i][j]=1
// 				// maxNum=1
// 			}
// 		}
// 	}


// 	for i := 1; i < len(text1Bytes); i++ {
// 		for j := 1; j < len(text2Bytes); j++ {
// 			if text1Bytes[i]==text2Bytes[j] {
// 				dp[i][j]=dp[i-1][j-1]+1
// 				// maxNum=max(maxNum,dp[i][j])
// 			}else{
// 				dp[i][j]=max(dp[i-1][j],dp[i][j-1])
// 			}
// 		}
// 	}

// 	return maxNum

// }

// 1035
// dp[i][j]表示上面到i，下面到j的最大连线数
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp:=make([][]int,len(nums1))
	for i := 0; i < len(nums1); i++ {
		dp[i]=make([]int, len(nums2))
	}
	
	// 初始化
	if nums1[0]==nums2[0] {
		dp[0][0]=1
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i]==nums2[0] {
			for j := i; j < len(nums1); j++ {
				dp[j][0]=1
			}
			break
		}
	}
	for i := 0; i < len(nums2); i++ {
		if nums1[0]==nums2[i] {
			for j := i; j < len(nums2); j++ {
				dp[0][j]=1
			}
			break
		}
	}



	for i := 1; i < len(nums1); i++ {
		for j := 1; j < len(nums2); j++ {
			if  nums1[i]==nums2[j]{
				dp[i][j]=max(max(dp[i-1][j-1]+1,dp[i][j-1]),dp[i-1][j])
			}else{
				dp[i][j]=max(max(dp[i-1][j-1],dp[i][j-1]),dp[i-1][j])
			}
		}
	}

	return dp[len(nums1)-1][len(nums2)-1]
}

// 53
func maxSubArray(nums []int) int {
	dp:=make([]int,len(nums))
	dp[0]=nums[0]
	res:=dp[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1]>0 {
			dp[i]=dp[i-1]+nums[i]
		}else{
			dp[i]=nums[i]
		}
		if res<dp[i] {
			res=dp[i]
		}
	}
	return res
}

// 392
// dp[i][j]表示是否在长度为j的序列t中存在长度为i的序列s
func isSubsequence(s string, t string) bool {
	if len(s)==0 {
		return true
	}
	if len(t)==0{
		return false
	}
	ss:=[]byte(s)
	ts:=[]byte(t)
	dp:=make([][]int,len(ss))
	for i := 0; i < len(ss); i++ {
		dp[i]=make([]int, len(ts))
	}

	for i := 0; i < len(ts); i++ {
		if ts[i]==ss[0] {
			for j := 0; j < len(ts); j++ {
				dp[0][j]=1
			}
			break
		}
	}


	for i := 1; i < len(ss); i++ {
		for j := 1; j < len(ts); j++ {
			if ss[i]==ts[j]{
				dp[i][j]=dp[i-1][j-1]+1
			}else{
				dp[i][j]=dp[i][j-1]
			}
		}
	}

	return dp[len(ss)-1][len(ts)-1]==len(ss)
}

// 115
// dp[i][j]表示i-1序列在t[0-j-1]中的重复次数
// dp[i][j]=dp[i-1][j-1]+1    不同：dp[i][j]=
// func numDistinct(s string, t string) int {

// }

// 583 && 72
func minDistance(word1 string, word2 string) int {
	dp:=make([][]int,len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i]=make([]int, len(word2)+1)
	}
	for i := 0; i < len(word1)+1; i++ {
		dp[i][0]=i
	}
	for i := 0; i < len(word2)+1; i++ {
		dp[0][i]=i
	}
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1]==word2[j-1] {
				dp[i][j]=dp[i-1][j-1]
			}else{
				dp[i][j]=min(dp[i-1][j]+1,min(dp[i][j-1]+1,dp[i-1][j-1]+2))   
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(a,b int)int{
	if a<b {
		return a
	}
	return b
}