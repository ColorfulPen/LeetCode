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

// 剑指 Offer 05. 替换空格 LCOF
func replaceSpace(s string) string {
	return nil
}

// 151. Reverse Words in a String
func reverseWords(s string) string {
    startIndex:=0
    for s[startIndex]==' '{
        startIndex++
    }
	bytePool:=make([]byte,0)
	res:=make([]byte,0)
	for i := len(s)-1; i >= startIndex; i-- {
		if s[i]==' ' {
			// 把pool里的内容移动，同时清空pool
			for i := len(bytePool)-1; i >= 0; i-- {
				res = append(res, bytePool[i])
				if i==0 {
					res = append(res, ' ')
				}
			}
			bytePool=make([]byte, 0)
		}else{
			bytePool = append(bytePool, s[i])
		}
	}
	for i := len(bytePool)-1; i >= 0; i-- {
		res = append(res, bytePool[i])

	}
	return string(res[:len(res)])
}

// 剑指 Offer 58 - II. 左旋转字符串 LCOF
func reverseLeftWords(s string, n int) string {
	reverseS :=func (s []byte,from,to int) string {
		reverseRange:=to-from
		for i := 0; i < reverseRange/2; i++ {
			s[from+i],s[to-1-i]=s[to-1-i],s[from+i]
		}
		return string(s)
	}
	n=n%len(s)
	s=reverseS([]byte(s),0,n)
	s=reverseS([]byte(s),n,len(s))
	s=reverseS([]byte(s),0,len(s))
	return s
}