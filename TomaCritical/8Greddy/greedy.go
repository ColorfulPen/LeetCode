package greddy

import "sort"

// 455. Assign Cookies
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	indexG, indexS := 0, 0
	for indexG < len(g) && indexS < len(s) {
		if s[indexS] >= g[indexG] {
			indexG++
			indexS++
			count++
		} else {
			indexS++
		}
	}
	return count
}

// 376. Wiggle Subsequence
func wiggleMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	count := 1
	startIndex := 1
	var flag bool
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			startIndex = i
			if nums[i] > nums[i-1] {
				count++
				flag = true
			} else if nums[i] < nums[i-1] {
				count++
				flag = false
			}
			break
		}
	}
	for i := startIndex + 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] && !flag {
			count++
			flag = true
		} else if nums[i] < nums[i-1] && flag {
			count++
			flag = false
		}
	}
	return count
}

// 122. Best Time to Buy and Sell Stock II
func maxProfit(prices []int) int {
	max := 0
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			max += profit
		}
	}
	return max
}

// 55. Jump Game
func canJump(nums []int) bool {
	farest := 0
	for i := 0; i <= farest; i++ {
		farest = max(farest, i+nums[i])
		if farest >= len(nums) {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 45. 跳跃游戏 II
func jump(nums []int) int {
	// 记录最远的距离，在这距离之间查找第二远的距离，并记录，每记录一次加一
	if len(nums)==1 {
		return 0
	}
	count := 1
	farest := nums[0]
	i := 0
	for farest < len(nums)-1 {
		maxDis := 0
		for i <= farest {
			temp := i + nums[i]
			maxDis = max(maxDis, temp)
			i++
		}
		count++
		farest = maxDis
		if farest >= len(nums)-1 {
			return count
		}
	}

	return count
}

// 1005. K 次取反后最大化的数组和
func largestSumAfterKNegations(nums []int, k int) int {
	specialSort:=func (nums *[]int)  {
		sort.Ints(*nums)
	}
	specialSort(&nums)
	index:=0
	for i := 0; i < k; i++ {
		if index==len(nums) {
			if (k-i)%2==1 {
				nums[index-1]=-nums[index-1]
			}
			return sum(nums)
		}
		if nums[index]<0 {
			nums[index]=-nums[index]
		}
		index++
	}
	return sum(nums)
}

func sum(nums []int) int{
	total:=0
	for i := 0; i < len(nums); i++ {
		total+=nums[i]
	}
	return total
}

// 134. 加油站
func canCompleteCircuit(gas []int, cost []int) int {
	rest:=0
	rightIndex:=len(cost)-1
	for i := 0; i <= rightIndex; i++ {
		rest+=gas[i]-cost[i]
		for rest<0 {
			rest+=gas[rightIndex]-cost[rightIndex]
			if rightIndex==i {
				return -1
			}
			rightIndex--
		}
	}
	return (rightIndex+1)%len(cost)
}

// 135. 分发糖果
func candy(ratings []int) int {
	count:=0
	candy:=make([]int,len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i]>ratings[i-1] {
			candy[i]=candy[i-1]+1
		}
	}
	for i := len(ratings)-2; i >=0 ; i-- {
		if ratings[i]>ratings[i+1] && candy[i]<=candy[i+1]{
            candy[i]=candy[i+1]+1
		}
	}
    for i:=0;i<len(ratings);i++{
        count+=candy[i]
    }
	return count+len(ratings)
}

// 860. 柠檬水找零
func lemonadeChange(bills []int) bool {
	rest:=make([]int,2)
	for i := 0; i < len(bills); i++ {
		if bills[i]==5 {
			rest[0]++
		}else if bills[i]==10{
			rest[0]--
			rest[1]++
			if rest[0]<0 {
				return false
			}
		}else{
			if rest[1]>0 {
				rest[1]--
				if rest[0]>0 {
					rest[0]--
				}else{
					return false
				}
			}else{
				if rest[0]>2 {
					rest[0]=rest[0]-3
				}else{
					return false
				}
			}
		}
	}
	return true
}


