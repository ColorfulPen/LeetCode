/*
 * @Author: TomaChen513
 * @Date: 2022-11-20 10:09:05
 * @LastEditors: TomaChen513
 * @LastEditTime: 2022-11-21 20:39:52
 * @FilePath: /LeetCode/CXYLeetCodeGo/GreddyAlo/greddy.go
 * @Description:
 *
 * Copyright (c) 2022 by TomaChen513 xy_chen513@163.com, All Rights Reserved.
 */

// 445
// 从小满足
// g: 1,3,5,7,9
// s: 5,5,5,7
package greddyalo

import "sort"

// 455
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	sIndex, gindex, count := 0, 0, 0
	for sIndex != len(s) || gindex != len(g) {
		if s[sIndex] < g[gindex] {
			sIndex++
		} else {
			sIndex++
			gindex++
			count++
		}
	}
	return count
}

// 376
// 我的算法太不精炼了
func wiggleMaxLength(nums []int) int {
	nums = removeDup(nums)
	numLen := len(nums)
	if numLen == 1 || numLen == 2 {
		return numLen
	}
	var prevState bool
	if nums[1] > nums[0] {
		prevState = true
	} else {
		prevState = false
	}
	for i := 2; i < len(nums); i++ {
		var currState bool
		if nums[i] > nums[i-1] {
			currState = true
		} else {
			currState = false
		}
		if currState == prevState {
			nums = removeEle(nums, i)
			i--
		} else {
			prevState = !prevState
		}
	}
	return len(nums)
}

func removeEle(nums []int, index int) []int {
	res := nums[:index]
	temp := nums[index+1:]
	res = append(res, temp...)
	return res
}

func removeDup(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			nums = removeEle(nums, i)
			i--
		}
	}
	return nums
}

func wiggleMaxLengthBetter(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	ans := 1
	preDiff := nums[1] - nums[0]
	if preDiff != 0 {
		ans = 2
	}
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		// diff==preDiff 说明坡度相同，不需要考虑
		if diff > 0 && preDiff <= 0 || diff < 0 && preDiff >= 0 {
			ans++
			preDiff = diff
		}
	}
	return ans
}

// 53
func maxSubArrary(nums []int) int {
	res := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if res < count {
			res = count
		}
		if count < 0 {
			count = 0
		}
	}
	return res
}

// 122
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	res := 0
	for i := 1; i < len(prices); i++ {
		money := prices[i] - prices[i-1]
		if money > 0 {
			res += money
		}
	}
	return res
}

// 55
// 贪心，每次按最远的跳，如果跳跃范围内的点的跳跃值加该点下标大于原先跳跃的点（最大的那个），则选择此点为落脚点
func canJump(nums []int) bool {
	index := 0
	for index < len(nums) {
		prev := index
		farestIndex := index + nums[index]

		if farestIndex >= len(nums)-1 {
			return true
		}

		nextIndex := index
		for j := index + 1; j <= farestIndex; j++ {
			subFarStep := j + nums[j]
			if subFarStep > farestIndex {
				farestIndex = subFarStep
				nextIndex = j
			}
		}

		if farestIndex >= len(nums)-1 {
			return true
		}

		index = nextIndex

		if prev == index {
			break
		}
	}
	return false
}

func canJumpSimp(nums []int) bool {
	cover:=0+nums[0]
	// 对只有一个元素时的判断
    if cover>=len(nums)-1 {
        return true
    }

	for i := 1; i <= cover; i++ {
		subCover:=i+nums[i]
		if subCover>=len(nums)-1 {
			return true
		}
		if subCover>cover {
			cover=subCover
		}
	}
	return false
}

// 45
func jump(nums []int) int {
	count:=0
	index := 0
	for index < len(nums) {
		prev := index
		farestIndex := index + nums[index]

		nextIndex := index
		for j := index + 1; j <= farestIndex; j++ {
			subFarStep := j + nums[j]
			if subFarStep > farestIndex {
				farestIndex = subFarStep
				nextIndex = j
			}
		}

		index = nextIndex

	}

	return count
}
