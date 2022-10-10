package cn

import (
	"testing"
)

func TestValidAnagram(t *testing.T) {
	t.Log(isAnagram("anagram", "nagaram"))
	t.Log(isAnagram("rat", "car"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	//var sMap,tMap map[byte]int
	sMap, tMap := make(map[byte]int, len(s)), make(map[byte]int, len(t))
	for i := 0; i < len(s); i++ {
		//a := s[i]
		sMap[s[i]]++
		tMap[t[i]]++
	}

	for k, v := range sMap {
		tv, ok := tMap[k]
		if !ok {
			return false
		}
		if tv != v {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
