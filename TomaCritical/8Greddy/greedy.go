package greddy

import "sort"

// 455. Assign Cookies
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	indexG, indexS := 0, 0
	for indexG < len(g) && indexS < len(s) {
		if s[indexS] >= g[indexG] {
			indexG++
			indexS++
			count++
		} else {
			indexS++
		}
	}
	return count
}

// 376. Wiggle Subsequence
func wiggleMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	count := 1
	startIndex := 1
	var flag bool
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			startIndex = i
			if nums[i] > nums[i-1] {
				count++
				flag = true
			} else if nums[i] < nums[i-1] {
				count++
				flag = false
			}
			break
		}
	}
	for i := startIndex+1; i < len(nums); i++ {
		if nums[i] > nums[i-1] && !flag {
			count++
			flag = true
		} else if nums[i] < nums[i-1] && flag {
			count++
			flag = false
		}
	}
	return count
}

