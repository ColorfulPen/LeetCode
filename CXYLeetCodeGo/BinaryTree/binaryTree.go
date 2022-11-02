package binarytree

import (
	"container/list"
	"math"
)

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

func levelTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}
	queue := list.New()
	res = make([]int, 0)
	queue.PushBack(root)
	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)
		node := e.Value.(*TreeNode)
		res = append(res, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return res
}

// 226
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// 101
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSym(root.Left, root.Right)
}

func isSym(leftNode *TreeNode, rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}
	return leftNode != nil && rightNode != nil && leftNode.Val == rightNode.Val && isSym(leftNode.Left, rightNode.Right) && isSym(leftNode.Right, rightNode.Left)
}

// 104
func maxDepth(root *TreeNode) int{
	if root==nil {
		return 0
	}
	leftHeight,rightHeight:=0,0
	if root.Left!=nil {
		leftHeight=maxDepth(root.Left)
	}
	if root.Right!=nil {
		rightHeight=maxDepth(root.Right)
	}
	return int(math.Max(float64(leftHeight),float64(rightHeight)))+1
}

// 111
func minDepth(root *TreeNode) int{
	if root==nil {
		return 0
	}
	if root.Left==nil {
		return minDepth(root.Right)+1
	}
	if root.Right==nil {
		return minDepth(root.Left)+1
	}
	return int(math.Min(float64(minDepth(root.Right)),float64(minDepth(root.Left))))+1
}

func countNodes(root *TreeNode) int{
	if root==nil {
		return 0
	}
	queue:=list.New()
	queue.PushBack(root)
	res:=0
	for queue.Len()!=0 {
		n:=queue.Len()
		for i := 0; i < n; i++ {
			node:=queue.Remove(queue.Front()).(*TreeNode)
			if node.Left!=nil {
				queue.PushBack(node.Left)
			}
			if node.Right!=nil {
				queue.PushBack(node.Right)
			}
			res++
		}
	}
	return res
}
