package DoubleElements

import(
	"testing"
)

func Test(t *testing.T){
	t.Log(removeElement([]int{3,2,2,3},3))
}

func removeElement(nums []int, val int) int {
	len:=len(nums)
	if len==0{
		return 0
	}
	leftIndex,rightIndex:=0,len-1
	for{
		if nums[leftIndex]!=val {
			leftIndex++
		} else {
			// 交换
			temp:=nums[leftIndex]
			nums[leftIndex]=nums[rightIndex]
			nums[rightIndex]=temp
			rightIndex--
		}
		if leftIndex>rightIndex{
			break
		}
	}
	return rightIndex+1
}