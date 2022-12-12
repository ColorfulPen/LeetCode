package stackandqueue

import (
	"container/list"
	"fmt"
)

// 20. Valid Parentheses
func isValid(s string) bool {
	stack:=list.New()
	sByte:=[]byte(s)
	sMap:=make(map[byte]byte,3)
	sMap['}']='{'
	sMap[']']='['
	sMap[')']='('
	for i := 0; i < len(s); i++ {
		if key,ok:=sMap[sByte[i]];ok{
			if stack.Len()==0 {
				return false
			}else{
				curr:=stack.Remove(stack.Back()).(byte)
				if curr!=key {
					return false
				}
			}
		}else{
			stack.PushBack(sByte[i])
		}
	}
	if stack.Len()==0 {
		return true
	}else{
		return false
	}
}

// 1047. Remove All Adjacent Duplicates In String
func removeDuplicates(s string) string {
	sBytes:=[]byte(s)
	stack:=list.New()
	for i := len(s)-1; i >= 0; i-- {
		if stack.Len()!=0 {
			temp:=stack.Back().Value.(byte)
			if temp==sBytes[i] {
				stack.Remove(stack.Back())
			}else{
				stack.PushBack(sBytes[i])
			}
		}else{
			stack.PushBack(sBytes[i])
		}
	}
	res:=make([]byte,0)
	length:=stack.Len()
	for i := 0; i < length; i++ {
		temp:=stack.Remove(stack.Back()).(byte)
		res = append(res, temp)
	}
	return string(res)
}