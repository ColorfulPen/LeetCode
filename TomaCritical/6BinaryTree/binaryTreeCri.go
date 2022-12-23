package binarytree

import "container/list"

// 102. Binary Tree Level Order Traversal
// 层序遍历队列版
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	queue := list.New()
	if root == nil {
		return res
	}
	queue.PushBack(root)
	length := 1
	for queue.Len() != 0 {
		levelRes := make([]int, 0)
		for length > 0 {
			curr := queue.Remove(queue.Front()).(*TreeNode)
			levelRes = append(levelRes, curr.Val)
			if curr.Left != nil {
				queue.PushBack(curr.Left)
			}
			if curr.Right != nil {
				queue.PushBack(curr.Right)
			}
			length--
		}
		res = append(res, levelRes)
		length = queue.Len()
	}
	return res
}

// 层序遍历递归版
// 知晓思路
func levelOrder_Recursion(root *TreeNode) [][]int {
	res:=make([][]int,0)
	if root==nil {
		return res
	}
	var recursion func(node *TreeNode,level int,res [][]int)
	recursion = func (node *TreeNode,level int,res [][]int){
		if len(res)<level {
			res = append(res, make([]int, 0))
		}
		res[level-1]=append(res[level-1], node.Val)
		if node.Left!=nil {
			recursion(node.Left,level+1,res)
		}
		if node.Right!=nil {
			recursion(node.Right,level+1,res)
		}
	}
	recursion(root,1,res)
	return res
}


