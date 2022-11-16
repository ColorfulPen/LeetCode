/*
 * @Author: TomaChen513
 * @Date: 2022-11-11 10:18:05
 * @LastEditors: TomaChen513
 * @LastEditTime: 2022-11-16 11:24:31
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