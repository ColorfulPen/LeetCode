package DoubleElements

import(
	"testing"
)

func Test(t *testing.T){
	t.Log(removeElement([]int{3,2,2,3},3))
	t.Log(replaceSpace("We are happy."))
}

// 27
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

// 344
func reverseString(s []byte){
	left,right:=0,len(s)-1
	// for{
	// 	// temp:=s[left]
	// 	// s[left]=s[right]
	// 	// s[right]=temp
	// 	s[left],s[right]=s[right],s[left]
	// 	left++
	// 	right--
	// 	if left>=right {
	// 		break
	// 	}
	// }

	// while
	for left<right{
		s[left],s[right]=s[right],s[left]
		left++
		right--
	}
}

//offer 05
func replaceSpace(s string) string {
	b:=[]byte(s)
	length:=len(b)
	spaceCount:=0
	for _,char :=range s{
		if(char==' '){
			spaceCount++
		}
	}
	tmp:=make([]byte,spaceCount*2)
	b=append(b,tmp...)

	i:=length-1
	j:=len(b)-1
	for i>=0{
		if b[i]!=' '{
			b[j]=b[i]
			i--
			j--
		}else{
			b[j],b[j-1],b[j-2]='0','2','%'
			i--
			j-=3
		}
	}
	return string(b)
}

type MyQueue struct{
	queue []int
}

func NewMyQueue() *MyQueue{
	return &MyQueue{
		queue: make([]int,0),
	}
}

func (m *MyQueue) Front() int{
	return m.queue[0];
}

func (m *MyQueue) Back() int{
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue) Empty() bool{
	return len(m.queue)==0
}

func (m *MyQueue) Push(val int){
	for !m.Empty() && m.Back()<val{
		m.queue=m.queue[:len(m.queue)-1]
	}
	m.queue=append(m.queue,val)
}

func (m *MyQueue) Pop(val int){
	if !m.Empty() && m.Front()==val{
		m.queue=m.queue[1:]
	}
}



// 239
func maxSlidingWindow(nums []int, k int) []int {
	queue:=NewMyQueue()
	length:=len(nums)
	res:=make([]int,0)

	// 初始化部分
	for i:=0;i<k;i++{
		queue.Push(nums[i])
	}

	res=append(res,queue.Front())

	for i:=k;i<length;i++{
		queue.Pop(nums[i-k])
		queue.Push(nums[i])
		res=append(res,queue.Front())
	}

	return res
}