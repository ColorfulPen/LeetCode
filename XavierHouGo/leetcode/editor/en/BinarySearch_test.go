package cn

import (
    "testing"
)

func TestBinarySearch(t *testing.T) {
	t.Log(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	t.Log(search([]int{0, 1}, 1))
	t.Log(search([]int{0, 1}, 0))
    t.Log(search([]int{0}, 0))

}

//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	var mid int
	for left < right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}
        if nums[mid] < target {
			left = mid + 1
            continue
		}
        if nums[mid] > target {
			right = mid
            continue
		}

	}

	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
