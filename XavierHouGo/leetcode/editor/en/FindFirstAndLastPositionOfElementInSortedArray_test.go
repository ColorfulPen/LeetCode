package cn

import (
	"testing"
)

func TestFindFirstAndLastPositionOfElementInSortedArray(t *testing.T) {
	t.Log(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	t.Log(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	t.Log(searchRange([]int{}, 0))
	t.Log(searchRange([]int{1}, 1))
}

//leetcode submit region begin(Prohibit modification and deletion)

func binarySearch(nums []int, target int, equal bool) int {
	left := 0
	right := len(nums)
	var mid int

	index := right
	for left < right {
		mid = (left + right) / 2

		if nums[mid] > target || (equal && nums[mid] >= target) {
			right = mid
			index = mid
			continue
		} else {
			left = mid + 1
		}
	}

	return index

}

func searchRange(nums []int, target int) []int {
	start := binarySearch(nums, target, true)
	end := binarySearch(nums, target, false) - 1
	if start > end || start == -1 || end == -1 {
		return []int{-1, -1}
	}
	return []int{start, end}
}

//leetcode submit region end(Prohibit modification and deletion)
