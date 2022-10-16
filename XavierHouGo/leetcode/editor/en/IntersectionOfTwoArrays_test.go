package cn

import (
    "testing"
)

func TestIntersectionOfTwoArrays(t *testing.T) {
    t.Log(intersection([]int{1,2,2,1},[]int{2,2}))
    t.Log(intersection([]int{4,9,5},[]int{9,4,9,8,4}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
    hash := make([]int,1002)
    res := make([]int,0)

    for _,v := range nums1 {
        hash[v] = 1
    }

    for _,v := range nums2 {
        if hash[v] == 1 {
            res = append(res, v)
            hash[v] = 2
        }
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)

