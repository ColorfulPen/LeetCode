package cn

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	t.Log(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
	k := 0

	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left] == val {
			for nums[right] == val && left < right {
				right--
				k++
			}
			nums[left], nums[right] = nums[right], nums[left]
			k++
			right--
		}
		left++
	}
	return len(nums) - k
}

//leetcode submit region end(Prohibit modification and deletion)
