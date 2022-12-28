package array

import "math"

// sort Algorithm
// https://blog.csdn.net/qq_51664685/article/details/124427443
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	i, j := start, end
	baseval := arr[i]
	for i < j {
		for i < j && arr[j] >= baseval {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] <= baseval {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = baseval

	quickSort(arr, start, i-1)
	quickSort(arr, i+1, end)
}

// 选择排序
// 选择排序算法的原理如下：

// ​ 1.首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置。

// ​ 2.再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。

// ​ 3.重复第二步，直到所有元素均排序完毕。

// 插入排序算法的原理如下：

// ​ 1.从第一个元素开始，该元素可以认为已经被排序；

// ​ 2.取出下一个元素，在已经排序的元素序列中从后向前扫描；

// ​ 3.如果该元素（已排序）大于新元素，将该元素移到下一位置；

// ​ 4.重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；

// ​ 5.将新元素插入到该位置后；

// ​ 6.重复步骤2~5。

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
			window := index - startIndex
			minWindow = min(window, minWindow)
			// move
			sum -= nums[startIndex]
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
