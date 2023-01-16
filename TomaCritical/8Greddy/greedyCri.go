package greddy

import "sort"

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

// 406.根据身高重建队列
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people,func(i, j int) bool {
		if people[i][0]==people[j][0] {
			return people[i][1]<people[j][1]
		}
		return people[i][0]>people[j][0]
	})
	for i := 0; i < len(people); i++ {
		copy(people[people[i][1]+1:],people[people[i][1]:])
		people[people[i][1]]=people[i]
	}
	return people
}