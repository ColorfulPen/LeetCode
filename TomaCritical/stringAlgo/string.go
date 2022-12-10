package stringalgo


// 344. Reverse String
func reverseString(s []byte)  {
	length:=len(s)
	for i := 0; i < length/2; i++ {
		s[i],s[length-1-i]=s[length-1-i],s[i]
	}
}

// 541. Reverse String II
func reverseStr(s string, k int) string {
	sbytes:=[]byte(s)
	reversePart:=func (s []byte,from,to int) []byte {
		length:=to-from
		for i := 0; i < length/2; i++ {
			s[i+from],s[to-1-i]=s[to-1-i],s[i+from]
		}
		return s
	}
	length:=len(s)
	index:=0
	for length-index>0 {
		left:=length-index
		if left>2*k {
			sbytes=reversePart(sbytes,index,index+k)
			index+=2*k
		}else if left>=k{
			sbytes=reversePart(sbytes,index,index+k)
			index+=left
		}else{
			sbytes=reversePart(sbytes,index,index+left)
			index+=left
		}
	}
	return string(sbytes)
}