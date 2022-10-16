package stringetc

import (
	"testing"
)

func TestReverse(t *testing.T) {
	t.Log(reverse("abcdef", 0, 2))
	t.Log(reverseLeftWords("abcdef", 3))
}

func TestStr(t *testing.T) {
	t.Log(strStr("hello", "ll"))
	t.Log(strStr("aaaaa", "baa"))

}



// 151. Reverse Words in a String
// 将输入的单词倒序，后整个翻转
// func reverseWords(s string) string {

// }

// 剑指 Offer 58 - II. 左旋转字符串
func reverseLeftWords(s string, n int) string {
	n = n % len(s)
	s = reverse(s, 0, n-1)
	s = reverse(s, n, len(s)-1)
	s = reverse(s, 0, len(s)-1)
	return s
}

// reverse word from specific index
func reverse(s string, leftIndex, rightIndex int) string {
	b := []byte(s)
	for leftIndex < rightIndex {
		// s[leftIndex],s[rightIndex]=s[rightIndex],s[leftIndex]
		b[leftIndex], b[rightIndex] = b[rightIndex], b[leftIndex]
		leftIndex++
		rightIndex--
	}
	return string(b)
}

// 28. Find the Index of the First Occurrence in a String
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := caluNext(needle)
	j := -1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j]
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == len(needle)-1 {
			return i - len(needle) + 1
		}
	}
	return -1
}

func caluNext(needle string) []int {
	res := make([]int, len(needle))
	j := -1
	res[0] = j
	for i := 1; i < len(needle); i++ {
		for j >= 0 && needle[i] != needle[j+1] {
			j = res[j]
		}
		if needle[i] == needle[j+1] {
			j++
		}
		res[i] = j
	}
	return res
}

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	next := make([]int, n)
	j := -1
	next[0] = j
	for i := 0; i < n; i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j
	}
	if next[n-1] != -1 && n%(n-(next[n-1]+1)) == 0 {
		return true
	}
	return false
}
