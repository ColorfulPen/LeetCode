package cn

import (
	"testing"
)

func TestFourSumIi(t *testing.T) {
	t.Log(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
	t.Log(fourSumCount([]int{0}, []int{0}, []int{0}, []int{0}))
	t.Log(fourSumCount([]int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	nums := 0
	n := len(nums1)
	aMap, bMap := make(map[int]int, n), make(map[int]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			aMap[nums1[i]+nums2[j]]++
		}

	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			bMap[nums3[i]+nums4[j]]++
		}

	}

	for k, av := range aMap {
		bv, ok := bMap[-k]
		if ok {
			nums = nums + av*bv
			delete(aMap, k)
			delete(bMap, -k)
		}
	}
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
