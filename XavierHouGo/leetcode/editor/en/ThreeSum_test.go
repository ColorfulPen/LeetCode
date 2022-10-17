package cn

import (
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	t.Log(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	t.Log(threeSum([]int{0, 1, 1}))
	t.Log(threeSum([]int{0, 0, 0}))
	t.Log(threeSum([]int{1, 2, -2, -1}))

}

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		target := -1 * nums[first]

		//for second := first+1;second<n;second++ {
		//    if second>first+1&&nums[second]==nums[second-1] {
		//        continue
		//    }
		//    for second<third&&nums[second] + nums[third] > target{
		//        third--
		//    }
		//    if second==third{
		//        break
		//    }
		//    if nums[second] + nums[third] == target {
		//        ans = append(ans, []int{nums[first],nums[second],nums[third]})
		//    }
		//}
		for second := first + 1; second < third; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
