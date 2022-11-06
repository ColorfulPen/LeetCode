package binarytree

import (
	"container/list"
	"math"
	"strconv"
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight, rightHeight := 0, 0
	if root.Left != nil {
		leftHeight = maxDepth(root.Left)
	}
	if root.Right != nil {
		rightHeight = maxDepth(root.Right)
	}
	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

// 111
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return int(math.Min(float64(minDepth(root.Right)), float64(minDepth(root.Left)))) + 1
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	res := 0
	for queue.Len() != 0 {
		n := queue.Len()
		for i := 0; i < n; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			res++
		}
	}
	return res
}

// 110
func isBalanced(root *TreeNode) bool {
	// 1.左右高度差小于1 2.左树and右树是平衡树
	if root == nil {
		return true
	}
	leftHeight, rightHeight := 0, 0
	if root.Left != nil {
		leftHeight = maxDepth(root.Left)
	}
	if root.Right != nil {
		rightHeight = maxDepth(root.Right)
	}
	abs := int(math.Abs(float64(leftHeight) - float64(rightHeight)))
	flag := true
	if abs > 1 {
		flag = false
	}
	return flag && isBalanced(root.Left) && isBalanced(root.Right)
}

// 257
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}
	traverse(&res, root, "")
	return res
}

func traverse(res *[]string, node *TreeNode, path string) {
	if node.Left == nil && node.Right == nil {
		path = path + strconv.Itoa(node.Val)
		*res = append(*res, path)
		return
	}
	path = path + strconv.Itoa(node.Val) + "->"
	if node.Left != nil {
		traverse(res, node.Left, path)
	}
	if node.Right != nil {
		traverse(res, node.Right, path)
	}
}

// 404
func sumOfLeftLeaves(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}
	return traverseLeft(root.Left, true) + traverseLeft(root.Right, false)

}

func traverseLeft(node *TreeNode, flag bool) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil && flag {
		return node.Val
	}
	return traverseLeft(node.Left, true) + traverseLeft(node.Right, false)
}

// 如果pop出最后一个，且这一个左右节点都是空，说明是最后一个
// 513
func findBottomLeftValue(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)
	var res int
	for queue.Len() > 0 {
		currNode := queue.Remove(queue.Front()).(*TreeNode)
		if currNode.Right != nil {
			queue.PushBack(currNode.Right)
		}
		if currNode.Left != nil {
			queue.PushBack(currNode.Left)
		}
		if queue.Len() == 0 {
			res = currNode.Val
		}
	}
	return res
}

// 112
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	subTargetSum := targetSum - root.Val
	if subTargetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, subTargetSum) || hasPathSum(root.Right, subTargetSum)

}

// 106
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 中序 LNR
	// 后序 LRN
	if len(inorder) == 0 {
		return nil
	}
	// var node *TreeNode
	// node.Val=postorder[len(postorder)-1]
	// inIndex:=findInIndex(inorder,node.Val)
	// node.Left=buildTree(inorder[:inIndex],postorder[:inIndex])
	// node.Right=buildTree(inorder[inIndex+1:],postorder[inIndex:len(postorder)-1])
	// 因为没有TreeNode的构造函数，所以会报错
	val := postorder[len(postorder)-1]
	inIndex := findInIndex(inorder, val)
	left := buildTree(inorder[:inIndex], postorder[:inIndex])
	right := buildTree(inorder[inIndex+1:], postorder[inIndex:len(postorder)-1])
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

func findInIndex(array []int, val int) int {
	res := -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			res = i
			break
		}
	}
	return res
}

func findMaxIndex(array []int) (int, int) {
	maxNum := array[0]
	maxNumIndex := 0
	for i := 1; i < len(array); i++ {
		if maxNum < array[i] {
			maxNum = array[i]
			maxNumIndex = i
		}
	}
	return maxNumIndex, maxNum
}

// 654
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	maxNumIndex, maxNum := findMaxIndex(nums)
	return &TreeNode{
		Val:   maxNum,
		Left:  constructMaximumBinaryTree(nums[:maxNumIndex]),
		Right: constructMaximumBinaryTree(nums[maxNumIndex+1:]),
	}
}

// 617
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1==nil{
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

// 700
func searchBST(root *TreeNode, val int) *TreeNode {
	if root==nil {
		return nil
	}
	if root.Val==val {
		return root
	}else if root.Val>val {
		return searchBST(root.Left,val)
	}else{
		return searchBST(root.Right,val)
	}
}

// 98
func isValidBST(root *TreeNode) bool {
	if root==nil {
		return true
	}
	if root.Left!=nil {
		if root.Val<root.Left.Val {
			return false
		}
	}
	if root.Right!=nil {
		if root.Val>root.Right.Val {
			return false
		}
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}


