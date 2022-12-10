package hashset

import "sort"

// 15. 3Sum
func threeSum(nums []int) [][]int {
	// 双指针法
	sort.Ints(nums)
	res:=make([][]int,0)
	for i := 0; i < len(nums); i++ {
		
		left,right:=i+1,len(nums)-1
		sum:=nums[i]+nums[left]+nums[right]
		for left!=right{
			if sum>0 {
				right--
				sum=sum+nums[right]-nums[right+1]
			}else if sum<0{
				left++
				sum=sum+nums[left]-nums[left-1]
			}else{
				res = append(res, []int{i,left,right})
			}
		}
	}
	return res
}

// 18. 4Sum
func fourSum(nums []int, target int) [][]int {
	// 和三数之和采用相同思路
	return nil
}