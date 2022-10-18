package stacks

import(
	"testing"
)


type MyQueue struct{
	inputStack []int
	outtputStack []int
}

func Constructor() MyQueue{
	return MyQueue{
		inputStack: make([]int, 0),
		outtputStack: make([]int, 0),
	}
}

func (queue *MyQueue) Push(x int){
	queue.inputStack=append(queue.inputStack, x)
}

func (queue *MyQueue) Pop() int{
	if len(queue.outtputStack)!=0 {
		res:=queue.outtputStack[len(queue.outtputStack)-1]
		queue.outtputStack=queue.outtputStack[:len(queue.outtputStack)-1]
		return res
	}

	for i:=len(queue.inputStack)-1;i>=1;i-- {
		queue.outtputStack = append(queue.outtputStack, queue.inputStack[i])
	}

	res:=queue.inputStack[0]
	queue.inputStack=queue.inputStack[:0]
	return res
}

func (queue *MyQueue) Peek() int{
	if len(queue.outtputStack)!=0 {
		return queue.outtputStack[len(queue.outtputStack)-1]
	}else{
		return queue.inputStack[0]
	}
}

func (queue *MyQueue) Empty() bool{
	if len(queue.inputStack)==0 && len(queue.outtputStack)==0 {
		return true
	}
	return false
}

func TestQueue(t *testing.T){
	testQueue:=Constructor()
	t.Log(testQueue.Empty())
	testQueue.Push(1)
	t.Log(testQueue.Empty())
	testQueue.Push(2)
	testQueue.Push(3)
	t.Log(testQueue.Pop())
	t.Log(testQueue.Pop())
}