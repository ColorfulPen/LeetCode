package binarytree

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NLR
func preorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

func preorderTraversalNotRecursion(root *TreeNode) (res []int) {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		ans = append(ans, node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return ans
}

func orderTraversalUnifiedIteration(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}
	stack := list.New()
	res = make([]int, 0)
	stack.PushBack(root)
	var node *TreeNode

	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)

		// 若为空标记节点
		if e.Value == nil {
			e = stack.Back()
			stack.Remove(e)
			node = e.Value.(*TreeNode)
			res = append(res, node.Val)
			continue
		}
		node = e.Value.(*TreeNode)
		if node.Right != nil {
			stack.PushBack(node)
		}
		if node.Left != nil {
			stack.PushBack(node)
		}
		stack.PushBack(node)
		stack.PushBack(nil)
	}
	return res
}

func levelTraversal(root *TreeNode) (res []int){
	if root==nil {
		return nil
	}
	queue:=list.New()
	res=make([]int, 0)
	queue.PushBack(root)
	for queue.Len()>0{
		e:=queue.Front()
		queue.Remove(e)
		node:=e.Value.(*TreeNode)
		res = append(res, node.Val)
		if node.Left!=nil {
			queue.PushBack(node.Left)
		}
		if node.Right!=nil {
			queue.PushBack(node.Right)
		}
	}
	return res
}
