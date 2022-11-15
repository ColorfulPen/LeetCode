/*
 * @Author: TomaChen513
 * @Date: 2022-11-11 09:39:43
 * @LastEditors: TomaChen513
 * @LastEditTime: 2022-11-15 10:23:48
 * @FilePath: /LeetCode/CXYLeetCodeGo/BackTracking/backTracking.go
 * @Description:
 *
 * Copyright (c) 2022 by TomaChen513 xy_chen513@163.com, All Rights Reserved.
 */
package backtracking

import (
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
    result:=make([][]int,0)
	cBack(n,k,1,[]int{},&result)
	return result
}

// 注意要引入一个新的path
func cBack(n,k,startIndex int,path []int,result *[][]int){
	if len(path)==k {
		// 注意这里拷贝的容量要足够
		temp:=make([]int,k)
		copy(temp,path)
		*result=append(*result, temp)
		return
	}
	
	for i := startIndex; i <= n-(k-len(path))+1; i++ {
		path = append(path, i)
		cBack(n,k,i+1,path,result)
		path=path[:len(path)-1]
	}
}

// 216
func combinationSum3(k int, n int) [][]int {
	result:=make([][]int,0)
	backTwo(k,n,1,0,[]int{},&result)
	return result
}

func backTwo(k,n,startIndex,sum int,path []int,result *[][]int){
	if len(path)==k{
		if sum==n {
			temp:=make([]int,k)
			copy(temp,path)
			*result=append(*result, temp)
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
		if sum+startIndex>n {
			return
		}
		backTwo(k,n,startIndex+1,sum+startIndex,path,result)
		path=path[:len(path)-1]
		startIndex++
	}
}

// 17
func letterCombinations(digits string) []string {
	if digits=="" {
		return nil
	}
	letterMap:=map[string]string{"2":"abc","3":"def","4":"ghi","5":"jkl","6":"mno","7":"pqrs","8":"tuv","9":"wxyz"}
	result:=make([]string,0)
	backThree(digits,"",letterMap,&result,1)
	return result
}

func backThree(digits,path string,letterMap map[string]string,result *[]string,startIndex int){
	if len(path)==len(digits) {
		temp:=strings.Clone(path)
		*result=append(*result, temp)
		return
	}

	for startIndex<=len(digits) {
		char:=digits[startIndex-1:startIndex]
		letters:=letterMap[char]
		for i := 0; i < len(letters); i++ {
			path+=letters[i:i+1]
			backThree(digits,path,letterMap,result,startIndex+1)
			path=path[:len(path)-1]
		}
		startIndex++
	}
}

// 39
func combinationSum(candidates []int, target int) [][]int {
	result:=make([][]int,0)
	backFour(candidates,[]int{},target,&result,0)
	return result
}

func backFour(candidates,path []int, target int,result *[][]int,startIndex int){
	if target<0 {
		return
	}
	if target==0 {
		temp:=make([]int,len(path))
		copy(temp,path)
		*result=append(*result, temp)
		return
	}

	for startIndex<len(candidates){
		target-=candidates[startIndex]
		path = append(path, candidates[startIndex])
		backFour(candidates,path,target,result,startIndex)
		path=path[:len(path)-1]
		target+=candidates[startIndex]
		startIndex++
	}
}

// 40
func combinationSum2(candidates []int, target int) [][]int {
	result:=make([][]int,0)
	sort.Ints(candidates)
	backFive(candidates,[]int{},target,&result,0)
	return result
}

func backFive(candidates,path []int,target int,result *[][]int,startIndex int){
	if target<0 {
		return
	}
	if target==0 {
		temp:=make([]int,len(path))
		copy(temp,path)
		*result=append(*result, temp)
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		if i>startIndex && candidates[i]==candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		target-=candidates[i]
		backFive(candidates,path,target,result,i+1)
		path=path[:len(path)-1]
		target+=candidates[i]
	}
}