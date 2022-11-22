/*
 * @Author: TomaChen513
 * @Date: 2022-11-22 13:46:41
 * @LastEditors: TomaChen513
 * @LastEditTime: 2022-11-22 13:52:48
 * @FilePath: /LeetCode/CXYLeetCodeGo/GreddyAlo/greddy_test.go
 * @Description:
 *
 * Copyright (c) 2022 by TomaChen513 xy_chen513@163.com, All Rights Reserved.
 */
package greddyalo

import "testing"

func Test134(t *testing.T) {
	canCompleteCircuit([]int{4,5,3,1,4},[]int{5,4,3,4,2})
	canCompleteCircuit([]int{1},[]int{1})

}