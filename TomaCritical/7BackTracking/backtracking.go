package backtracking

import (
	"sort"
)

// 77. 组合
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var backtrack func(n, k, startIndex int, path []int)
	backtrack = func(n, k, startIndex int, path []int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := startIndex; i <= n-k+len(path)+1; i++ {
			path = append(path, i)
			backtrack(n, k, i+1, path)
			path = path[:len(path)-1]
		}
	}
	backtrack(n, k, 1, []int{})
	return res
}

// 216. Combination Sum III
func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)

	var back func(k, n, startIndex, total int, path []int)
	back = func(k, n, startIndex, total int, path []int) {
		if len(path) == k {
			if total == n {
				temp := make([]int, k)
				copy(temp, path)
				result = append(result, temp)
			}
			return
		}
		if total >= n {
			return
		}
		for i := startIndex; i < 10-k+len(path)+1; i++ {
			path = append(path, i)
			total += i
			back(k, n, i+1, total, path)
			total -= i
			path = path[:len(path)-1]
		}
	}

	back(k, n, 1, 0, []int{})
	return result
}

// 17. Letter Combinations of a Phone Number
func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	letterMap := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	result := make([]string, 0)
	var back func(digits string, letterMap map[byte]string, path []byte, startIndex int)
	back = func(digits string, letterMap map[byte]string, path []byte, startIndex int) {
		if len(path) == len(digits) {
			tmp := string(path)
			result = append(result, tmp)
			return
		}

		letters := letterMap[digits[startIndex]]
		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i])
			back(digits, letterMap, path, startIndex+1)
			path = path[:len(path)-1]
		}

	}
	back(digits, letterMap, []byte{}, 0)
	return result
}

// 39. Combination Sum
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	var back func(cancandidates []int, target, startIndex int, path []int)
	back = func(cancandidates []int, total, startIndex int, path []int) {
		if total == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		if total > target {
			return
		}

		// 注意边界条件
		for i := startIndex; i < len(cancandidates); i++ {
			path = append(path, cancandidates[i])
			total += cancandidates[i]
			back(cancandidates, total, i, path)
			path = path[:len(path)-1]
			total -= cancandidates[i]
		}

	}
	back(candidates, 0, 0, []int{})
	return result
}

// 40. 组合总和 II
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	var back func(candidates []int, target, total, startIndex int, path []int)
	back = func(candidates []int, target, total, startIndex int, path []int) {
		if total == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		if total > target {
			return
		}

		for i := startIndex; i < len(candidates); i++ {
			//
			if i != startIndex && candidates[i] == candidates[i-1] {
				continue
			}
			total += candidates[i]
			path = append(path, candidates[i])
			back(candidates, target, total, i+1, path)
			total -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	back(candidates, target, 0, 0, []int{})
	return res
}

// 131. 分割回文串
func partition(s string) [][]string {
	result := make([][]string, 0)
	isPalid := func(str string, from, to int) bool {
		length := to - from
		for i := from; i <= from+length/2; i++ {
			if str[i] != str[to+from-1-i] {
				return false
			}
		}
		return true
	}
	var back func(s string, startIndex int, path []string)
	back = func(s string, startIndex int, path []string) {
		if startIndex == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		// 具体的回溯，画图计算会更好一些
		for i := startIndex; i < len(s); i++ {
			if isPalid(s, startIndex, i+1) {
				path = append(path, s[startIndex:i+1])
				back(s, i+1, path)
				path = path[:len(path)-1]
			}
		}

	}

	back(s, 0, []string{})
	return result
}

// 93. 复原 IP 地址
func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	isValid := func(s string, from, to int) bool {
		length := to - from
		if (length > 1 && s[from] == '0') || length > 3 {
			return false
		}
		num := 0
		for i := from; i < to; i++ {
			num *= 10
			num += int(s[i] - '0')
		}
		if num > 255 {
			return false
		}
		return true
	}
	var back func(s string, startIndex int, path []string)
	back = func(s string, startIndex int, path []string) {
		if startIndex == len(s) && len(path) == 4 {
			temp := path[0] + "." + path[1] + "." + path[2] + "." + path[3]
			result = append(result, temp)
			return
		}

		for i := startIndex; i < len(s); i++ {
			if isValid(s, startIndex, i+1) {
				path = append(path, s[startIndex:i+1])
				back(s, i+1, path)
				path = path[:len(path)-1]
			}
		}
	}
	back(s, 0, []string{})
	return result
}

// 78. 子集
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	var back func(nums []int, startIndex int, path []int)
	back = func(nums []int, startIndex int, path []int) {
		// if startIndex==len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)
		// 	return
		// }

		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			back(nums, i+1, path)
			path = path[:len(path)-1]
		}
	}
	back(nums,0,[]int{})
	return result
}

// 90.子集II
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result:=make([][]int,0)
	var back func(nums []int,startIndex int,path []int)
	back=func(nums []int, startIndex int, path []int) {
		// add to result
		temp:=make([]int,len(path))
		copy(temp,path)
		result = append(result, temp)

		for i := startIndex; i < len(nums); i++ {
			if i>startIndex && nums[i]==nums[i-1] {
				continue
			}			
			path = append(path, nums[i])
			back(nums,i+1,path)
			path=path[:len(path)-1]
		}
	}
	back(nums,0,[]int{})
	return result
}

// 491. 递增子序列
func findSubsequences(nums []int) [][]int {
	result:=make([][]int,0)
	var back func(nums []int,startIndex int,path []int)
	back=func(nums []int, startIndex int, path []int) {
		if len(path)>=2 {
			temp:=make([]int,len(path))
			copy(temp,path)
			result = append(result, temp)
		}

		numSet:=make(map[int]struct{},0)
		for i := startIndex; i < len(nums); i++ {
			// 因为没有排序所以不能用
			// if i>startIndex && nums[i]==nums[i-1] {
			// 	continue
			// }
			if _,ok:=numSet[nums[i]];ok {
				continue
			}
			if len(path)==0 || path[len(path)-1]<=nums[i] {
				numSet[nums[i]]=struct{}{}
				path=append(path, nums[i])
				back(nums,i+1,path)
				path=path[:len(path)-1]
			}
		}
	}
	back(nums,0,[]int{})
	return result
}

// 46. 全排列
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	numMap:=make(map[int]struct{},0)
	var back func(nums,path []int,numMap map[int]struct{})
	back=func(nums, path []int,numMap map[int]struct{}) {
		if len(path)==len(nums) {
			temp:=make([]int,len(path))
			copy(temp,path)
			result = append(result, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if _,ok:=numMap[nums[i]];ok {
				continue
			}
			path = append(path, nums[i])
			numMap[nums[i]]=struct{}{}
			back(nums,path,numMap)
			path=path[:len(path)-1]
			delete(numMap,nums[i])
		}
	}
	back(nums,[]int{},numMap)
	return result
}

// 47. 全排列 II
// 注意这是层之间与深度之间的去重
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	indexMap:=make(map[int]struct{},0)
	var back func(nums,path []int,indexMap map[int]struct{})
	back=func(nums, path []int,indexMap map[int]struct{}) {
		if len(path)==len(nums) {
			temp:=make([]int,len(path))
			copy(temp,path)
			result = append(result, temp)
			return
		}

		levelUsed:=make(map[int]struct{},0)
		for i := 0; i < len(nums); i++ {
			if _,ok:=indexMap[i];ok {
				continue
			}
			if _,ok:=levelUsed[nums[i]];ok {
				continue
			}

			path = append(path, nums[i])
			indexMap[i]=struct{}{}
			levelUsed[nums[i]]=struct{}{}
			back(nums,path,indexMap)
			path=path[:len(path)-1]
			delete(indexMap,i)
		}
	}
	back(nums,[]int{},indexMap)
	return result
}
