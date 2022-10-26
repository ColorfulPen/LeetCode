package cn

import (
	"testing"
)

func TestFruitIntoBaskets(t *testing.T) {
	t.Log(totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func totalFruit(fruits []int) int {
	max := 0
	left := 0
	right := 0

	ht := make(map[int]int)
	for right < len(fruits) {
		if len(ht) <= 2 {
			ht[fruits[right]]++
			right++
			if len(ht) <= 2 {
				cur := right - left
				if cur > max {
					//fmt.Println(ht)
					max = cur
				}
			}

		}

		if len(ht) > 2 {
			ht[fruits[left]]--
			if ht[fruits[left]] == 0 {
				delete(ht, fruits[left])
			}
			left++
		}

	}
	//fmt.Println(ht)
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
