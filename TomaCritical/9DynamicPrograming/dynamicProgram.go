package dynamicprograming

// 53. Maximum Subarray
func maxSubArray(nums []int) int {
	res := nums[0]
	dp:=make([]int,len(nums))
	dp[0]=nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i]=max(dp[i-1]+nums[i],nums[i])
		res=max(dp[i],res)
	}
	return res
}

func max(a,b int)int{
	if a>b {
		return a
	}
	return b
}

// 70爬楼梯
func climbStairs(n int) int {
	if n==1 {
		return 1
	}
	dp:=make([]int,n+1)
	dp[1]=1
	dp[2]=2
	for i := 3; i < n+1; i++ {
		dp[i]=dp[i-1]+dp[i-2]
	}
	return dp[n]
}

// 746. 使用最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
	dp:=make([]int,len(cost)+1)
	if len(cost)==1{
		return 0
	}
	for i := 2; i < len(cost)+1; i++ {
		dp[i]=min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])	
	}

	return dp[len(cost)]
}

func min(a,b int)int{
	if a<b {
		return a	
	}
	return b
}

// 62. 不同路径  63. 不同路径 II
func uniquePaths(m int, n int) int {
	dp:=make([][]int,m)
	for i := 0; i < m; i++ {
		dp[i]=make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0]=1
	}
	for i := 0; i < n; i++ {
		dp[0][i]=1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j]=dp[i][j-1]+dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}