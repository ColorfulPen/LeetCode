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