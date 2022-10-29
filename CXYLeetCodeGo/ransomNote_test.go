package main

import (
	"testing"
)

func TestRansom(t *testing.T) {

	t.Log(canConstruct("xcx", "asd"))

}

func canConstruct(ransomNote string, magazine string) bool {
	// var chars [26]int

	chars := make([]int, 26)

	for _, c := range magazine {
		chars[c-'a']++
	}

	// for i:=0;i<len(magazine);i++{
	// 	chars[magazine[i]-'a']++
	// }
	for i := 0; i < len(ransomNote); i++ {
		chars[ransomNote[i]-'a']--
		if chars[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
