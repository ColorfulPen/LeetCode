package binarytree

import (
	"math"
	"strconv"

	
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 226. Invert Binary Tree
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	return root
}

// 101. Symmetric Tree
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return subIsSymmetric(root.Left, root.Right)
}

func subIsSymmetric(node1, node2 *TreeNode) bool {
	if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) {
		return false
	} else if node1 == nil && node2 == nil {
		return true
	}
	return node1.Val == node2.Val && subIsSymmetric(node1.Left, node2.Right) && subIsSymmetric(node1.Right, node2.Left)
}

// 104. Maximum Depth of Binary Tree
func maxDepth(root *TreeNode) int {
	return subMaxDepth(root, 0)
}

func subMaxDepth(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}
	left := subMaxDepth(node.Left, depth+1)
	right := subMaxDepth(node.Right, depth+1)
	if left > right {
		return left
	} else {
		return right
	}
}

// 111. Minimum Depth of Binary Tree
func minDepth(root *TreeNode) int {
	return subMinDepth(root, 0)
}

func subMinDepth(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}
	if node.Left == nil && node.Right == nil {
		return depth + 1
	}
	var left, right int
	if node.Left != nil {
		left = subMinDepth(node.Left, depth+1)
	} else {
		left = 100000
	}
	if node.Right != nil {
		right = subMinDepth(node.Right, depth+1)
	} else {
		right = 100000
	}
	if left < right {
		return left
	} else {
		return right
	}
}

// 222. 完全二叉树的节点个数
func countNodes(root *TreeNode) int {
	isComplete := func(node *TreeNode) int {
		// assume not null
		left, right := node.Left, node.Right
		height := 1
		for left != nil && right != nil {
			left = left.Left
			right = right.Right
			height++
		}
		if left == nil && right == nil {
			return int(math.Pow(2, float64(height)) - 1)
		}
		return -1
	}

	var recursiveCount func(node *TreeNode) int
	recursiveCount = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		totalNum := isComplete(node)
		if totalNum != -1 {
			return totalNum
		}
		return 1 + recursiveCount(node.Left) + recursiveCount(node.Right)
	}
	return recursiveCount(root)
}

// 110. 平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight, rightHeight := countHeight(root.Left, 1), countHeight(root.Right, 1)

	if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func countHeight(node *TreeNode, height int) int {
	if node == nil {
		return height - 1
	}
	leftHeight := countHeight(node.Left, height+1)
	rightHeight := countHeight(node.Right, height+1)
	if leftHeight > rightHeight {
		return leftHeight
	}
	return rightHeight
}

// 257. 二叉树的所有路径
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}
	var back func(node *TreeNode, path string)
	back = func(node *TreeNode, path string) {
		if node.Left == nil && node.Right == nil {
			path += strconv.Itoa(node.Val)
			res = append(res, path)
			return
		}
		path = path + strconv.Itoa(node.Val)
		if node.Left!=nil {
			back(node.Left, path+"->")
		}
		if node.Right!=nil {
			back(node.Right,path+"->")			
		}
	}
	back(root,"")
	return res
}

// 404. 左叶子之和
func sumOfLeftLeaves(root *TreeNode) int {
	if root==nil {
		return 0
	}
	var handle func(node *TreeNode,left bool) int
	handle=func(node *TreeNode,left bool) int {
		if node==nil {
			return 0
		}
		if node.Left==nil && node.Right==nil{
			if left {
				return node.Val
			}else{
				return 0
			}
		}
		return handle(node.Left,true)+handle(node.Right,false)
	}
	return handle(root.Left,true)+handle(root.Right,false)
}

// 513. 找树左下角的值
func findBottomLeftValue(root *TreeNode) int {
	res:=root.Val
	maxDepth:=1
	var dfs func(node *TreeNode,depth int)
	dfs=func(node *TreeNode, depth int) {
		if node.Left==nil && node.Right==nil {
			if depth>=maxDepth {
				res=node.Val
				maxDepth=depth
			}
			return
		}
		if node.Right!=nil {
			dfs(node.Right,depth+1)
		}
		if node.Left!=nil {
			dfs(node.Left,depth+1)
		}
	}
	dfs(root,1)
	return res
}

// 112. 路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root==nil {
		return false
	}
	count:=0
	var dfs func(node *TreeNode,sum int)
	dfs=func(node *TreeNode, sum int) {
		sum+=node.Val
		if node.Left==nil&&node.Right==nil {
			if targetSum==sum {
				count++
			}
		}
		if node.Left!=nil {
			dfs(node.Left,sum)
		}
		if node.Right!=nil {
			dfs(node.Right,sum)
		}
	}
	dfs(root,0)
	return count>0
}

// 106. 从中序与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	// LNR  LRN
	findMid:=func (inorder []int,mid int) int {
		for i := 0; i < len(inorder); i++ {
			if inorder[i]==mid {
				return i
			}
		}
		return -1
	}
	var buildTrees func(inorder,postorder []int) *TreeNode
	buildTrees=func(inorder, postorder []int) *TreeNode {
		if len(inorder)==1 {
			return &TreeNode{Val: inorder[0]}
		}
		midNum:=postorder[len(postorder)-1]
		node:=TreeNode{Val: midNum}
		index:=findMid(inorder,midNum)
		if index!=0 {
			node.Left=buildTree(inorder[:index],postorder[:index])
		}
		if  len(inorder)-index!=1{
			node.Right=buildTree(inorder[index+1:],postorder[index:len(postorder)-1])
		}
		return &node
	}
	return buildTrees(inorder,postorder)
}

// 654. 最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	findMaxIndex:=func (nums []int) (int,int) {
		max,res:=-1,-1
		for i := 0; i < len(nums); i++ {
			if nums[i]>max {
				max=nums[i]
				res=i
			}
		}
		return max,res
	}
	var buildMaxBinaryTree func(nums []int)*TreeNode
	buildMaxBinaryTree=func(nums []int) *TreeNode {
		max,index:=findMaxIndex(nums)
		node:=TreeNode{Val: max}
		if index>0 {
			node.Left=buildMaxBinaryTree(nums[:index])
		}
		if len(nums)-index>1 {
			node.Right=buildMaxBinaryTree(nums[index+1:])
		}
		return &node
	}
	return buildMaxBinaryTree(nums)
}

// 617. 合并二叉树
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1==nil {
		return root2
	}
	if root2==nil {
		return root1
	}
	root1.Val+=root2.Val
	root1.Left=mergeTrees(root1.Left,root2.Left)
	root1.Right=mergeTrees(root1.Right,root2.Right)
	return root1
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root==nil {
		return nil
	}
	if root.Val==val {
		return root
	}else if root.Val>val {
		return searchBST(root.Left,val)
	}else {
		return searchBST(root.Right,val)
	}
}

// 98. 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	return subIsValidBST(root,math.MinInt,math.MaxInt)
}

func subIsValidBST(root *TreeNode,min,max int) bool{
	if root==nil {
		return true
	}
	if root.Left!=nil && (root.Val<=root.Left.Val || min>=root.Left.Val){
		return false
	}
	if root.Right!=nil && (root.Val>=root.Right.Val || max<=root.Right.Val){
		return false
	}
	return subIsValidBST(root.Left,min,root.Val)&&subIsValidBST(root.Right,root.Val,max)
}

// 530. 二叉搜索树的最小绝对差
func getMinimumDifference(root *TreeNode) int {
	arr:=make([]int,0)
	var getSubDiff func(root *TreeNode)
	getSubDiff=func (root *TreeNode)  {
		if root==nil {
			return
		}
		if root.Left!=nil {
			getSubDiff(root.Left)
		}
		arr = append(arr, root.Val)
		if root.Right!=nil {
			getSubDiff(root.Right)
		}
	}
	getSubDiff(root)
	min:=math.MaxInt
	for i := 0; i < len(arr)-1; i++ {
		diff:=arr[i+1]-arr[i]
		if diff<min {
			min=diff
		}
	}
	return min
}

