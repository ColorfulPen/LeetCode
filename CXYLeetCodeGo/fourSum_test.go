package CXYLeetCodeGo

import(
	"testing"
)

func TestFourSum(t *testing.T){
	t.Log(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count:=0
	len:=len(nums1)
	aMap:=make(map[int]int)

	for i:=0;i<len;i++ {
		for j:=0;j<len;j++{
			aMap[nums1[i]+nums2[j]]++
		}
	}

	// for i:=0;i<len;i++{
	// 	for j:=0;j<len;j++{
	// 		count+=aMap[-nums3[i]-nums4[j]]
	// 	}
	// }
	
	// enchanced for loop
	for _,v3 :=range nums3{
		for _,v4 :=range nums4{
			count+=aMap[-v3-v4]
		}
	}

	return count;
}