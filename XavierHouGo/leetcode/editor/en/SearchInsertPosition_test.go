package cn

import (
    "testing"
)

func TestSearchInsertPosition(t *testing.T) {
    t.Log(searchInsert([]int{1,3,5,6},5))
    t.Log(searchInsert([]int{1,3,5,6},2))
    t.Log(searchInsert([]int{1,3,5,6},7))
    t.Log(searchInsert([]int{1,3,5,6},-1))
}

//leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
    for i:=0;i<len(nums);i++ {
        if nums[i] >= target {
            return i
        }
    }
    return len(nums)
}
//leetcode submit region end(Prohibit modification and deletion)

