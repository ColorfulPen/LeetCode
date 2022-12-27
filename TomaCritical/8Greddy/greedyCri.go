package greddy

// 53. Maximum Subarray
func maxSubArray(nums []int) int {
	res := nums[0]
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > res {
			res = count
		}
		if count < 0 {
			count = 0
			continue
		}

	}
	return res
}
