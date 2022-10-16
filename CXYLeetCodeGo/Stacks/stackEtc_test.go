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

func isValid(s string) bool{
	parenthesis:=map[byte]byte{')':'(',']':'[','}':'{'}
	stack:=make([]byte,0)
	if s==""{
		return true
	}
	for _,c :=range s{
		if c=='('||c=='{'||c=='['{
			stack=append(stack,byte(c))
		}else if len(stack)>0 && stack[len(stack)-1]==parenthesis[byte(c)]{
			stack=stack[:len(stack)-1]
		}else{
			return false
		}
	}

	return len(stack)==0
}
