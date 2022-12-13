package array

import "math"

// 704 Binary Search
func search(nums []int, target int) int {
	len := len(nums)
	left, right := 0, len
	for left != right {
		mid := (right - left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// 27 Remove Element
func removeElement(nums []int, val int) int {
	len := len(nums) - 1
	for i := 0; i < len; i++ {
		if nums[i] == val {
			nums[i], nums[len] = nums[len], nums[i]
			len--
			i--
		}
	}
	return len
}

// 977 Squares of a Sorted Array
func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	res := make([]int, 0)
	for left <= right {
		// 若左边大
		if math.Abs(float64(nums[left])) > math.Abs(float64(nums[right])) {
			res = append(res, nums[left]*nums[left])
			left++
		} else {
			res = append(res, nums[right]*nums[right])
			right--
		}
	}
	for i := 0; i < len(nums)/2; i++ {
		res[i], res[len(nums)-1-i] = res[len(nums)-1-i], res[i]
	}
	return res
}

// 209 Minimum Size Subarray Sum
func minSubArrayLen(target int, nums []int) int {
	// 加一个尾巴的0，来计算到最后一个数字
	nums = append(nums, 0)
	//滑动窗口
	minWindow := math.MaxInt
	sum := 0
	startIndex := 0
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for index := 0; index < len(nums); {
		if sum < target {
			sum += nums[index]
			index++
		} else {
			// record
			window:=index-startIndex
			minWindow=min(window,minWindow)
			// move
			sum-=nums[startIndex]
			startIndex++
		}
	}
	if minWindow == math.MaxInt {
		return 0
	}
	return minWindow
}

// 59 Spiral Matrix II
func generateMatrix(n int) [][]int {
	// init
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	mid := n / 2
	// 分奇偶填充中间值
	if n%2 != 0 {
		res[mid][mid] = n * n
	}
	count := 1
	for i := 0; i < n/2; i++ {
		length := n - 1 - i - i
		for j := i; j < n-1-i; j++ {
			res[i][j] = count
			res[j][n-1-i] = count + length
			res[n-1-i][n-1-j] = count + length*2
			res[n-1-j][i] = count + length*3
			count++
		}
		count = count + length*3
	}
	return res
}
