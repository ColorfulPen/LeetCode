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

// 501. 二叉搜索树中的众数
func findMode(root *TreeNode) []int {
	count,maxCount:=0,1
	res:=make([]int,0)
	var pre *TreeNode
	pre=nil
	var traverse func(node *TreeNode)

	traverse=func(node *TreeNode) {
		if node==nil {
			return
		}
		traverse(node.Left)

		if pre==nil {
			// res = append(res, node.Val)
			count=1
		}else if pre.Val==node.Val {
			count++
		}else{
			// if count>maxCount {
			// 	maxCount=count
			// 	res=res[:0]
			// 	res = append(res, pre.Val)
			// }else if count==maxCount {
			// 	res=append(res, pre.Val)
			// }
			count=1
		}

		if count==maxCount {
			res = append(res, node.Val)
		}
		if count>maxCount {
			res=res[:0]
			maxCount=count
			res = append(res, node.Val)
		}
		pre=node

		traverse(node.Right)
	}

	traverse(root)
	return res
}


