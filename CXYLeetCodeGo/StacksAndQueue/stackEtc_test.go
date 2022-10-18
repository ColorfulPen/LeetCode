package stacks

import (
	"container/list"
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

func maxSlidingWindow(nums []int, k int) []int {
	increQueue := list.New()
	increQueue.PushBack(nums[0])
	for i := 0; i < k; i++ {
		ele:=increQueue.Front()
		if ele==nil || int(*ele.Value)==nums[i] {

		}
	}
}
