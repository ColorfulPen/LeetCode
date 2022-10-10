package cn

import (
    "testing"
)

func change(arr []int) {
    arr[0] = 2
}

func TestTwoSum(t *testing.T) {
   arr := []int{0,1,2}
   change(arr)
   t.Log(arr)
}

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
    hashT := map[int]int{}

    for i,x := range nums {
        // 找key是target-x是否存在，若存在v是对应的下标
        if v,ok := hashT[target-x];ok {
            return []int{v,i}
        }
        // key为是nums的值，value是nums下标
        hashT[x] = i
    }
    return []int{}
}
//leetcode submit region end(Prohibit modification and deletion)

