package stackandqueue

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
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

// 150. Evaluate Reverse Polish Notation
func evalRPN(tokens []string) int {
	stack:=list.New()
	opMap:=make(map[byte]int8)
	opMap['+']=1
	opMap['-']=2
	opMap['*']=3
	opMap['/']=4
	for i := 0; i < len(tokens); i++ {
		if value,ok:=opMap[tokens[i][0]];ok {
			num1,_:=strconv.ParseInt(stack.Remove(stack.Back()).(string),10,64)
			num2,_:=strconv.ParseInt(stack.Remove(stack.Back()).(string),10,64)
			if value==1 {
				ans:=num1+num2
				stack.PushBack(strconv.FormatInt(ans,10))
			}else if value==2{
				stack.PushBack(strconv.FormatInt(num2-num1,10))
			}else if value==3{
				stack.PushBack(strconv.FormatInt(num1*num2,10))
			}else{
				stack.PushBack(strconv.FormatInt(num2/num1,10))
			}
		}else{
			stack.PushBack(tokens[i])
		}
	}
	res,_:=strconv.ParseInt(stack.Back().Value.(string),10,64)
	return int(res)
}

