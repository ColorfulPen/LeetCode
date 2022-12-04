/*
 * @Author: TomaChen513
 * @Date: 2022-11-30 10:19:08
 * @LastEditors: TomaChen513
 * @LastEditTime: 2022-11-30 15:14:02
 * @FilePath: /LeetCode/CXYLeetCodeGo/DynamicProgramming/dynamicProgram_test.go
 * @Description:
 *
 * Copyright (c) 2022 by TomaChen513 xy_chen513@163.com, All Rights Reserved.
 */
package dynamicprogramming

import "testing"

func Test300(t *testing.T) {
	lengthOfLIS([]int{1,3,6,7,9,4,10,5,6})
}

func Test718(t *testing.T){
	findLength([]int{1,2,3,2,1},[]int{3,2,1,4,7})
}

func Test1035(t *testing.T){
	maxUncrossedLines([]int{1,4,2},[]int{1,2,4})
}