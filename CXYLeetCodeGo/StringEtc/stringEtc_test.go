package stringetc

import(
	"testing"
)

func Test(t *testing.T){
	t.Log(reverse("abcdef",0,2))
	t.Log(reverseLeftWords("abcdef",3))
}

// 151. Reverse Words in a String
// 将输入的单词倒序，后整个翻转
// func reverseWords(s string) string {

// }

// 剑指 Offer 58 - II. 左旋转字符串
func reverseLeftWords(s string, n int) string {
	n=n%len(s)
	s=reverse(s,0,n-1)
	s=reverse(s,n,len(s)-1)
	s=reverse(s,0,len(s)-1)
	return s
}

// reverse word from specific index
func reverse(s string,leftIndex ,rightIndex int) string{
	b:=[]byte(s)
	for leftIndex<rightIndex{
		// s[leftIndex],s[rightIndex]=s[rightIndex],s[leftIndex]
		b[leftIndex],b[rightIndex]=b[rightIndex],b[leftIndex]
		leftIndex++
		rightIndex--
	}
	return string(b)
}



