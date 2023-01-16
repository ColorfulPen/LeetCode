package singlestack

import (
	"fmt"
	"strconv"
)

func main() {
	height := []int{5, 0, 2, 1, 4, 0, 1, 0, 3}
	Beans1 := CaluBeansDoublePoints(height)
	Beans2 := CaluBeansDynamicPrograming(height)

	fmt.Println("使用双指针法最大豆子数为：" + strconv.Itoa(Beans1))
	fmt.Println("使用双指针法最大豆子数为：" + strconv.Itoa(Beans2))
}

// 42. 接雨水
// 双指针法：从两端依次遍历到中央
// 思路：每列所能装的最多豆子取决于该列左和右的最大列高
func CaluBeansDoublePoints(height []int) int {
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0
	maxBeans := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				maxBeans += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				maxBeans += maxRight - height[right]
			}
			right--
		}
	}
	return maxBeans
}

// 动态规划：依次记录最大的左值和最大右值
func CaluBeansDynamicPrograming(height []int) int {
	maxBeans := 0
	leftMaxHeight := make([]int, len(height))
	rightMaxHeight := make([]int, len(height))
	leftMaxHeight[0] = height[0]
	rightMaxHeight[len(height)-1] = height[len(height)-1]

	// calu maxium left
	for i := 1; i < len(height); i++ {
		if height[i] > leftMaxHeight[i-1] {
			leftMaxHeight[i] = height[i]
		} else {
			leftMaxHeight[i] = leftMaxHeight[i-1]
		}
	}

	// calu maxium right
	for i := len(height) - 2; i >= 0; i-- {
		if height[i] > rightMaxHeight[i+1] {
			rightMaxHeight[i] = height[i]
		} else {
			rightMaxHeight[i] = rightMaxHeight[i+1]
		}
	}

	// calu all beans
	for i := 1; i < len(height)-1; i++ {
		h := -height[i]
		if leftMaxHeight[i] < rightMaxHeight[i] {
			h += leftMaxHeight[i]
		} else {
			h += rightMaxHeight[i]
		}
		if h > 0 {
			maxBeans += h
		}
	}
	return maxBeans
}
