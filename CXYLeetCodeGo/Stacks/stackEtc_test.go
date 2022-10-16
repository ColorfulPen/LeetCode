package stacks

import(
	"testing"
)

func Test(t *testing.T)  {
	val:=[3]int{1,2,3}
	t.Log(len(val))
	val1:=val[:len(val)-1]
	t.Log(len(val1))
	
}
