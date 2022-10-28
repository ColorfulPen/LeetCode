package stacks

import (
	"testing"
	
)

func Test(t *testing.T) {
	val := [3]int{1, 2, 3}
	t.Log(len(val))
	val1 := val[:len(val)-1]
	t.Log(len(val1))
}

func isValid(s string) bool {
	parenthesis := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0)
	if s == "" {
		return true
	}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, byte(c))
		} else if len(stack) > 0 && stack[len(stack)-1] == parenthesis[byte(c)] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func removeDuplicates(s string) string {
	stack := make([]byte, 0)
	stack = append(stack, s[0])
	for i := 1; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else {
			if stack[len(stack)-1] == s[i] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		}

	}
	return string(stack)
}

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, v := range tokens {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			x1, x2 := stack[len(stack)-2], stack[len(stack)-1]
			switch v {
			case "+":
				stack[len(stack)-2] = x1 + x2
			case "-":
				stack[len(stack)-2] = x1 - x2
			case "*":
				stack[len(stack)-2] = x1 * x2
			case "/":
				stack[len(stack)-2] = x1 / x2
			}
			stack = stack[:len(stack)-1]
		} else {
			sum, rate := 0, 1
			for i := len(v) - 1; i >= 1; i-- {
				sum += int(v[i]-'0') * rate
				rate *= 10
			}
			if v[0] == '-' {
				sum = -sum
			} else {
				sum += int(v[0]-'0') * rate
			}
			stack = append(stack, sum)
		}
	}
	return stack[0]
}

func maxSlidingWindowTwo(nums []int, k int) []int {
	increQueue := list.New()
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i >= k && increQueue.Front().Value.(int) == nums[i-k] {
			increQueue.Remove(increQueue.Front())
		}
		if increQueue.Front() == nil || nums[i] > increQueue.Front().Value.(int) {
			increQueue.PushBack(nums[i])
		}
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	increQueue := list.New()
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		ele := increQueue.Front()
		if ele == nil || ele.Value.(int) > nums[i] {
			increQueue.PushBack(nums[i])
		} else {
			for increQueue.Len() > 0 {
				increQueue.Remove(increQueue.Front())
			}
			increQueue.PushBack(nums[i])
		}
	}
	res = append(res, increQueue.Front().Value.(int))

	for i := k; i < len(nums); i++ {
		ele := increQueue.Front()
		if ele.Value.(int) == nums[i-k] {
			increQueue.Remove(increQueue.Front())
			ele = increQueue.Front()
		}
		if ele == nil {
			increQueue.PushBack(nums[i])
		} else {
			for nums[i] > ele.Value.(int) {
				increQueue.Remove(increQueue.Front())
				ele = increQueue.Front()
			}
			increQueue.PushBack(nums[i])
		}
		res = append(res, increQueue.Front().Value.(int))
	}
	return res
}

func TestList(t *testing.T) {
	t.Log(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}


