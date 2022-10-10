package cn

import (
	"testing"
)

func TestSqrtx11(t *testing.T) {
	t.Log(mySqrt(4))
	t.Log(mySqrt(0))

	t.Log(mySqrt(8))
	t.Log(mySqrt(1))
	t.Log(mySqrt(9))

}

//leetcode submit region begin(Prohibit modification and deletion)
func isSqrt(i, x int) uint8 {
	if (i*i) <= x && ((i+1)*(i+1)) > x {
		return 0
	}
	if (i * i) > x {
		return 1
	}
	if (i+1)*(i+1) <= x {
		return 2
	}
	return 3
}

func mySqrt(x int) int {
	left := 0
	right := x + 1

	var mid int
	for left < right {
		mid = (left + right) / 2

		switch isSqrt(mid, x) {
		case 0:
			return mid
		case 1:
			right = mid
		case 2:
			left = mid + 1
		}
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
