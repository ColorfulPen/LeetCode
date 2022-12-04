package array

// 704
func search(nums []int, target int) int {
	len:=len(nums)
	left,right:=0,len
	for left!=right{
		mid:=(right-left)/2
		if nums[mid]==target {
			return mid
		}else if nums[mid]>target {
			right=mid-1
		}else{
			left=mid+1
		}
	}
	return -1
}