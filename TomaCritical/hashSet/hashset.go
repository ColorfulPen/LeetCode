package hashset

// 242
func isAnagram(s string, t string) bool {
    if len(s)!=len(t){
        return false
    }
	letterMap:=make(map[byte]int)
	for i := 0; i < len(s); i++ {
		letterMap[s[i]]+=1
	}
	for i := 0; i < len(t); i++ {
		if letterMap[t[i]]-=1;letterMap[t[i]]<0 {
			return false
		}
	}
	return true
}

// 349
func intersection(nums1 []int, nums2 []int) []int {
	res:=make([]int,0)
	numMap:=make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		numMap[nums1[i]]+=1
	}
	for i := 0; i < len(nums2); i++ {
		if _,ok:=numMap[nums2[i]];ok {
			delete(numMap,nums2[i])
			res = append(res, nums2[i])
		}
	}
	return res
}