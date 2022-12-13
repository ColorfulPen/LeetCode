package stackandqueue

// 239. Sliding Window Maximum
func maxSlidingWindow(nums []int, k int) []int {
	stack:=make([]int,0)
	res:=make([]int,0)
	// init
	for i := 0; i < k; i++ {
		if len(stack)==0 {
			stack = append(stack, nums[i])
		}else{
			if nums[i]>=stack[len(stack)-1] {
				stack = append(stack, nums[i])
			}
		}
	}
	res = append(res, stack[len(stack)-1])
	for i := k; i < len(nums); i++ {
		abandon:=nums[i-k]
		if abandon==stack[len(stack)-1] {
			stack=stack[:len(stack)-1]
		}
		if len(stack)==0 {
			stack = append(stack, nums[i])
		}else{
			if stack[len(stack)-1]<=nums[i] {
				stack=stack[:0]
				stack = append(stack, nums[i])
			}
		}
		res = append(res, stack[len(stack)-1])
	}
	return res
}