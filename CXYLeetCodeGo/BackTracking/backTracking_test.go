/*
 * @Author: TomaChen513
 * @Date: 2022-11-11 10:18:05
 * @LastEditors: TomaChen513
 * @LastEditTime: 2022-11-18 09:38:56
 * @FilePath: /LeetCode/CXYLeetCodeGo/BackTracking/backTracking_test.go
 * @Description: 
 * 
 * Copyright (c) 2022 by TomaChen513 xy_chen513@163.com, All Rights Reserved. 
 */
package backtracking

import (
	"testing"
	"fmt"
)

func Test77(t *testing.T) {
	combine(4,2)
}

func TestLegal(t *testing.T){
	
	fmt.Println(isLegal("123"))
	fmt.Println(isLegal("13"))
	fmt.Println(isLegal("1"))
	fmt.Println(isLegal("01"))
	fmt.Println(isLegal("0"))
	fmt.Println(isLegal("257"))


}

func TestIpAdd(t *testing.T){
	fmt.Println(restoreIpAddresses("0000"))
}

func TestSubSe(t *testing.T){
	findSubsequences([]int{4,6,7,7})
}

func Test47(t *testing.T){
	ans:=permuteUnique([]int{1,1,2})
	fmt.Println(ans)
}

func Test46(t *testing.T){
	permute([]int{1,2,3})
}

func Test332(t *testing.T){
	// findItinerary([][]string{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}})
	findItinerary([][]string{{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"}})
}
