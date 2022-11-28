/*
 * @Author: TomaChen513
 * @Date: 2022-11-20 10:09:05
 * @LastEditors: TomaChen513
 * @LastEditTime: 2022-11-28 10:12:10
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

import (
	"fmt"
	"sort"
	"strconv"
)

// 455
func findContentChildren(g []int, s []int) int {
	fmt.Println("findContentChildren")
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
	cover := 0 + nums[0]
	// 对只有一个元素时的判断
	if cover >= len(nums)-1 {
		return true
	}

	for i := 1; i <= cover; i++ {
		subCover := i + nums[i]
		if subCover >= len(nums)-1 {
			return true
		}
		if subCover > cover {
			cover = subCover
		}
	}
	return false
}

// 45
// 选择下一个能跳最远的位置
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	count := 1
	index := 0
	step := nums[0]

	if step >= len(nums)-1 {
		return count
	}
	for index < len(nums) {
		farestIndex := step
		nextIndex := index

		for j := index + 1; j <= step; j++ {
			subFarStep := j + nums[j]
			if subFarStep > farestIndex {
				farestIndex = subFarStep
				nextIndex = j
			}
		}

		index = nextIndex
		step = farestIndex
		count++
		if step >= len(nums)-1 {
			break
		}
	}

	return count
}

// 1005
// 排序后找到正负分界点
// Better: 按照绝对值正负进行排序
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	// 全正
	if nums[0] >= 0 {
		if k%2 == 1 {
			nums[0] = -nums[0]
		}
		return sumInts(nums)
	}

	index := 0
	for k > 0 {
		if nums[index] < 0 {
			nums[index] = -nums[index]
			index++
			// 若下标大于数组长度
			if index == len(nums) {
				if k%2 == 0 {
					nums[len(nums)-1] = -nums[len(nums)-1]
				}
				return sumInts(nums)
			}
		} else {
			if nums[index] < nums[index-1] {
				if k%2 == 1 {
					nums[index] = -nums[index]
				}
				return sumInts(nums)
			} else {
				if k%2 == 1 {
					nums[index-1] = -nums[index-1]
				}
				return sumInts(nums)
			}
		}
		k--
	}
	return sumInts(nums)
}

func sumInts(nums []int) int {
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}

// 134
func canCompleteCircuit(gas []int, cost []int) int {
	sum := gas[0]-cost[0]
	startIndex := 0
	currIndex := 1
	for currIndex%len(gas) != startIndex {
		if sum >= 0 {
			sum+=gas[currIndex]-cost[currIndex]
			currIndex++
		} else {
			if startIndex==0 {
				startIndex=len(gas)-1
			}else{
				startIndex--
			}
			sum+=gas[startIndex]-cost[startIndex]
		}
	}

	if sum<0 {
		return -1
	}
	return startIndex
}

// 135
// 从左到右和从右向左遍历
// 左到右为右比左大的情况    右到左为左比右大的情况
func candy(ratings []int) int {
	candys:=make([]int,len(ratings))
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i]<ratings[i+1] {
			candys[i+1]=candys[i]+1
		}
	}
	for i := len(ratings)-1; i >= 1; i-- {
		if ratings[i-1]>ratings[i] {
			if candys[i-1]<=candys[i] {
				candys[i-1]=candys[i]+1
			}
		}
	}
	res:=0
	for _,v:=range candys{
		res+=v
	}
	return res+len(ratings)
}

// 860
// 能找10就不找5
func lemonadeChange(bills []int) bool {
	pocket:=make([]int,2)
	for i := 0; i < len(bills); i++ {
		money:=bills[i]
		switch money {
		case 5:
			pocket[0]++
		case 10:
			pocket[0]--
			if pocket[0]==-1 {
				return false
			}
			pocket[1]++
		case 20:
			if pocket[1]>=1 && pocket[0]>=1 {
				pocket[1]--
				pocket[0]--
			}else if pocket[0]>=3 {
				pocket[0]=pocket[0]-3
			}else{
				return false
			}
		}
	}
	return true
}

// 406
// 从高到低按身高排序  
// people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
// [7,0] [7,1] [6,1] [5,2] [5,0] [4,4] 
// [[7,0],[7,1],[6,1],[5,0],[5,2],[4,4]]

func reconstructQueue(people [][]int) [][]int {
	people=orderByHeight(people)
	for i := 0; i < len(people)-1; i++ {
		currPeople:=people[i]
		insertIndex:=people[i][1]
		// 全体后移，移到i处
		for j := i; j >insertIndex ; j-- {
			people[j]=people[j-1]
		}
		people[insertIndex]=currPeople
	}
	return people
}

// 5 2 1 4 3
// 冒泡排序
func orderByHeight(people [][]int) [][]int{
	for i := 0; i < len(people)-1; i++ {
		for j := 0; j < len(people)-1-i; j++ {
			if people[j][0]<people[j+1][0]{
				people[j],people[j+1]=people[j+1],people[j]
			}else if people[j][0]==people[j+1][0] &&  people[j][1]>people[j+1][1]{
				people[j],people[j+1]=people[j+1],people[j]
			}
		}
	}
	return people
}

func reconstructQueueBetter(people [][]int) [][]int {
	sort.Slice(people,func(i, j int) bool {
		if people[i][0]==people[j][0] {
			return people[i][1]<people[j][1]
		}
		return people[i][0]>people[j][0]
	})
	result:=make([][]int,0)

	for _,info:=range people{
		result=append(result, info)
		copy(result[info[1]+1:],result[info[1]:])
		result[info[1]]=info
	}
	return result
}

// 452
func findMinArrowShots(points [][]int) int {
	// 需要排序
	sort.Slice(points,func(i, j int) bool {
		return points[i][0]<points[j][0]
	})
	count:=1
	maxLeft,minRight:=points[0][0],points[0][1]
	for i := 1; i < len(points); i++ {
		left,right:=points[i][0],points[i][1]
		if minRight<left {
			count++
			maxLeft,minRight=left,right
		}else{
			if left>maxLeft {
				maxLeft=left
			}
			if right<minRight {
				minRight=right
			}
		}
	}
	return count
}

// 435
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals,func(i, j int) bool {
		if intervals[i][0]==intervals[j][0] {
			return intervals[i][1]<intervals[j][1]
		}
		return intervals[i][0]<intervals[j][0]
	})
	fmt.Println(intervals)
	count:=0
	preRight,preLeft:=intervals[0][1],intervals[0][0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0]<preRight && intervals[i][0]>preLeft{
			count++

		}else if intervals[i][0]==preLeft && intervals[i][1]==preRight{
			count++
		}else{
			preRight=intervals[i][1]
			preLeft=intervals[i][0]
		}
	}
	return count
}

// 763
func partitionLabels(s string) []int {
	res:=make([]int,0)
	farestIndex:=-1
	startIndex:=-1
	sMap:=make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		sMap[s[i]]=i
	}
	for i := 0; i < len(s); i++ {
		if sMap[s[i]]>farestIndex {
			farestIndex=sMap[s[i]]
		}
		if i==farestIndex {
			res = append(res, i-startIndex)
			startIndex=i
		}
	}
	return res
}

// 56
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals,func(i, j int) bool {
		return intervals[i][0]<intervals[j][0] 
	})
	res:=make([][]int,0)
	left,maxRight:=intervals[0][0],intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0]<=maxRight {
			if intervals[i][1]>maxRight {
				maxRight=intervals[i][1]
			}
		}else{
			res = append(res, []int{left,maxRight})
			left,maxRight=intervals[i][0],intervals[i][1]
		}
	}
	res = append(res, []int{left,maxRight})
	return res
}

// 738
func monotoneIncreasingDigits(n int) int {
	// 逆序数组
	divideNum:=make([]int,0)
	for n>0 {
		divideNum = append(divideNum, n%10)
		n/=10
	}
	fmt.Println(divideNum)
	for i := 0; i < len(divideNum)-1; i++ {
		if divideNum[i]<divideNum[i+1] {
			divideNum[i+1]-=1

			// 这里注意的是 只要前一位小于1，那么后面全部置9
			for j := 0; j <= i; j++ {
				divideNum[j]=9
			}
		}
	}
	res:=0
	rate:=1
	for i := 0; i < len(divideNum); i++ {
		res+=divideNum[i]*rate
		rate*=10
	}
	return res
}

func mods(N int)int{
	s:=strconv.Itoa(N)
	ss:=[]byte(s)
	n:=len(ss)
	if n<1 {
		return n
	}
	res,_:=strconv.Atoi(string(ss))
	return res
}

// 714
func maxProfit(prices []int, fee int) int {
	res:=0
	minPrice:=prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]<minPrice {
			minPrice=prices[i]
		}

		if prices[i]>=minPrice && prices[i]<minPrice+fee {
			continue
		}

		if prices[i]>minPrice+fee {
			res+=prices[i]-minPrice-fee
			minPrice=prices[i]-fee
		}
	}
	return res
}