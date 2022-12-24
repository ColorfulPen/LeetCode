package backtracking

// 77. 组合
func combine(n int, k int) [][]int {
	res:=make([][]int,0)
	var backtrack func(n,k,startIndex int,path []int)
	backtrack=func (n,k,startIndex int,path []int)  {
		if len(path)==k {
			temp:=make([]int,k)
			copy(temp,path)
			res = append(res, temp)
			return
		}
		for i := startIndex; i <= n-k+len(path)+1; i++ {
			path = append(path, i)
			backtrack(n,k,i+1,path)
			path=path[:len(path)-1]
		}
	}
	backtrack(n,k,1,[]int{})
	return res
}

// 216. Combination Sum III
func combinationSum3(k int, n int) [][]int {
	result:=make([][]int,0)

	var back func(k,n,startIndex,total int,path []int)
	back=func(k, n, startIndex, total int, path []int) {
		if len(path)==k {
			if total==n {
				temp:=make([]int,k)
				copy(temp,path)
				result = append(result, temp)
			}
			return
		}
		if total>=n {
			return
		}
		for i := startIndex; i < 10-k+len(path)+1; i++ {
			path = append(path, i)
			total+=i
			back(k,n,i+1,total,path)
			total-=i
			path=path[:len(path)-1]
		}
	}

	back(k,n,1,0,[]int{})
	return result
}


