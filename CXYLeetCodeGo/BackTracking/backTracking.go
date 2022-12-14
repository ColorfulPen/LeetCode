/*
 * @Author: TomaChen513
 * @Date: 2022-11-11 09:39:43
 * @LastEditors: TomaChen513
 * @LastEditTime: 2022-11-18 09:42:43
 * @FilePath: /LeetCode/CXYLeetCodeGo/BackTracking/backTracking.go
 * @Description:
 *
 * Copyright (c) 2022 by TomaChen513 xy_chen513@163.com, All Rights Reserved.
 */
package backtracking

import (
	"fmt"
	"sort"
	"strings"
)

// ------------
// func backTracking(){
// 	if endCondition {
// 		storeResult()
// 		return
// 	}
// 	for select this level element {
// 		handle this element
// 		backTracking()
// 		回溯
// 	}
// }
// ------------

// 77
func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	cBack(n, k, 1, []int{}, &result)
	return result
}

// 注意要引入一个新的path
func cBack(n, k, startIndex int, path []int, result *[][]int) {
	if len(path) == k {
		// 注意这里拷贝的容量要足够
		temp := make([]int, k)
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	for i := startIndex; i <= n-(k-len(path))+1; i++ {
		path = append(path, i)
		cBack(n, k, i+1, path, result)
		path = path[:len(path)-1]
	}
}

// 216
func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	backTwo(k, n, 1, 0, []int{}, &result)
	return result
}

func backTwo(k, n, startIndex, sum int, path []int, result *[][]int) {
	if len(path) == k {
		if sum == n {
			temp := make([]int, k)
			copy(temp, path)
			*result = append(*result, temp)
		}
		return
	}

	// for i := startIndex; i <= 9; i++ {
	// 	path = append(path, i)
	// 	// 这里的startIndex传入为i+1，原因是在一轮迭代过程中本质是由i决定的
	// 	backTwo(k,n,i+1,sum+i,path,result)
	// 	path=path[:len(path)-1]
	// }

	// 用这种表示更容易理解一点
	for startIndex <= 9 {
		path = append(path, startIndex)
		if sum+startIndex > n {
			return
		}
		backTwo(k, n, startIndex+1, sum+startIndex, path, result)
		path = path[:len(path)-1]
		startIndex++
	}
}

// 17
func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	letterMap := map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}
	result := make([]string, 0)
	backThree(digits, "", letterMap, &result, 1)
	return result
}

func backThree(digits, path string, letterMap map[string]string, result *[]string, startIndex int) {
	if len(path) == len(digits) {
		temp := strings.Clone(path)
		*result = append(*result, temp)
		return
	}

	for startIndex <= len(digits) {
		char := digits[startIndex-1 : startIndex]
		letters := letterMap[char]
		for i := 0; i < len(letters); i++ {
			path += letters[i : i+1]
			backThree(digits, path, letterMap, result, startIndex+1)
			path = path[:len(path)-1]
		}
		startIndex++
	}
}

// 39
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	backFour(candidates, []int{}, target, &result, 0)
	return result
}

func backFour(candidates, path []int, target int, result *[][]int, startIndex int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	for startIndex < len(candidates) {
		target -= candidates[startIndex]
		path = append(path, candidates[startIndex])
		backFour(candidates, path, target, result, startIndex)
		path = path[:len(path)-1]
		target += candidates[startIndex]
		startIndex++
	}
}

// 40
func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)
	backFive(candidates, []int{}, target, &result, 0)
	return result
}

func backFive(candidates, path []int, target int, result *[][]int, startIndex int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		target -= candidates[i]
		backFive(candidates, path, target, result, i+1)
		path = path[:len(path)-1]
		target += candidates[i]
	}
}

// 131
func partition(s string) [][]string {
	result := make([][]string, 0)
	backSix(s, []string{}, &result, 0, 0)
	return result
}

func backSix(s string, path []string, result *[][]string, startIndex, endIndex int) {
	if endIndex == len(s) && len(path) != 0 {
		temp := make([]string, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	for i := endIndex; i <= len(s); i++ {
		currSubString := s[startIndex:i]
		if !isPalid(currSubString) {
			continue
		}
		path = append(path, currSubString)
		backSix(s, path, result, i, i)
		path = path[:len(path)-1]
	}
}

func isPalid(subString string) bool {
	if subString == "" {
		return false
	}
	length := len(subString)
	byteArray := []byte(subString)
	for i := 0; i < length/2; i++ {
		if byteArray[i] != byteArray[length-i-1] {
			return false
		}
	}
	return true
}

// 93
func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	backSeven(s, []string{}, &result, 0, 0, 0)
	return result
}

func backSeven(s string, path []string, result *[]string, startIndex, endIndex, count int) {
	if endIndex == len(s) {
		if len(path) != 4 {
			return
		}
		ip := ""
		for i := 0; i < 4; i++ {
			ip += path[i]
			if i != 3 {
				ip += "."
			}
		}
		*result = append(*result, ip)
		return
	}

	for i := endIndex; i <= len(s); i++ {
		currSubString := s[startIndex:i]
		if !isLegal(currSubString) {
			continue
		}
		count++
		if count > 4 {
			return
		}
		path = append(path, currSubString)
		backSeven(s, path, result, i, i, count)
		count--
		path = path[:len(path)-1]
	}
}

func isLegal(subString string) bool {
	if subString == "" {
		return false
	}
	bytes := []byte(subString)
	if bytes[0] == '0' && len(subString) > 1 {
		return false
	}
	sum := 0
	for i := 0; i < len(subString); i++ {
		sum = sum*10 + int(bytes[i]-'0')
	}
	return sum < 255
}

// 78
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	backEight(nums, []int{}, &result, 0)
	return result
}

func backEight(nums, path []int, result *[][]int, startIndex int) {

	temp := make([]int, len(path))
	copy(temp, path)
	*result = append(*result, temp)

	if startIndex == len(nums) {
		return
	}

	for i := startIndex; i < len(nums); i++ {
		path = append(path, nums[i])
		backEight(nums, path, result, i+1)
		path = path[:len(path)-1]
	}
}

// 90
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)
	backEightDup(nums, []int{}, &result, 0)
	return result
}

func backEightDup(nums, path []int, result *[][]int, startIndex int) {

	temp := make([]int, len(path))
	copy(temp, path)
	*result = append(*result, temp)

	if startIndex == len(nums) {
		return
	}

	for i := startIndex; i < len(nums); i++ {
		if i != startIndex && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		backEightDup(nums, path, result, i+1)
		path = path[:len(path)-1]
	}
}

// 491
func findSubsequences(nums []int) [][]int {
	result := make([][]int, 0)
	backNight(nums, []int{}, &result, 0)
	return result
}

func backNight(nums, path []int, result *[][]int, startIndex int) {
	if len(path) >= 2 {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
	}

	if startIndex == len(nums) {
		return
	}

	uniMap := make(map[int]struct{}, 0)
	for i := startIndex; i < len(nums); i++ {
		// 因为是未排序的，所以不能单纯的通过前后关系去重
		// if i!=startIndex && nums[i]==nums[i-1] {
		// 	continue
		// }
		if len(path) >= 1 && nums[i] < path[len(path)-1] {
			continue
		}
		if _, ok := uniMap[nums[i]]; ok {
			continue
		}
		uniMap[nums[i]] = struct{}{}
		path = append(path, nums[i])
		backNight(nums, path, result, i+1)
		path = path[:len(path)-1]
	}
}

// 46
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	used := make([]bool, len(nums))
	backTen(nums, []int{}, &result, used)
	return result
}

func backTen(nums, path []int, result *[][]int, used []bool) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		backTen(nums, path, result, used)
		used[i] = false
		path = path[:len(path)-1]
	}
}

// 47
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)
	used := make([]bool, len(nums))
	backTen2(nums, []int{}, &result, used)
	return result
}

func backTen2(nums, path []int, result *[][]int, used []bool) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	// 在循环外控制的是层级间的重复
	paraUsed := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		// 带入递归算法内的控制的是纵向间的重复
		if used[i] {
			continue
		}
		if _, ok := paraUsed[nums[i]]; ok {
			continue
		}
		paraUsed[nums[i]] = struct{}{}
		used[i] = true
		path = append(path, nums[i])
		backTen2(nums, path, result, used)
		used[i] = false
		path = path[:len(path)-1]
	}
}

// 332
func findItinerary(tickets [][]string) []string {
	result := make([]string, 0)
	used := make([]bool, len(tickets))
	path := make([][]string, 0)
	backEleven(tickets, path, &result, used)
	return result
}

func backEleven(tickets [][]string, path [][]string, result *[]string, used []bool) {
	if len(path) == len(tickets) {
		temp:=make([]string,0)
		for i := 0; i < len(tickets); i++ {
			temp = append(temp, tickets[i][0])
		}
		temp = append(temp, tickets[len(tickets)][1])
		*result = temp
		fmt.Println(path)
	}

	prev := "ZZZ"
	for i := 0; i < len(tickets); i++ {

		if len(path) == 0 && tickets[i][0] != "JFK" {
			continue
		}
		if used[i] {
			continue
		}
		if len(path)!=0 && tickets[i][0]!=path[len(path)-1][1] {
			continue
		}
		if compareString(prev, tickets[i][1]) {
			continue
		}
		used[i] = true
		path = append(path, tickets[i])
		prev = tickets[i][1]
		if len(path) == len(tickets) {
			return
		}
		backEleven(tickets, path, result, used)
		used[i] = false
		path = path[:len(path)-1]
	}
}

// 第一个比第二个字典序小时，返回true
func compareString(dst1, dst2 string) bool {
	dst1Byte := []byte(dst1)
	dst2Byte := []byte(dst2)
	for i := 0; i < 3; i++ {
		if dst1Byte[i] < dst2Byte[i] {
			return true
		} else if dst1Byte[i] > dst2Byte[i] {
			return false
		}
	}
	return true
}
