package binarytree

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
	return subIsSymmetric(root.Left,root.Right)
}

func subIsSymmetric(node1,node2 *TreeNode) bool{
	if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) {
		return false
	} else if node1 == nil && node2 == nil {
		return true
	}
	return node1.Val==node2.Val && subIsSymmetric(node1.Left,node2.Right)  &&subIsSymmetric(node1.Right,node2.Left)
}

// 104. Maximum Depth of Binary Tree
func maxDepth(root *TreeNode) int {
	return subMaxDepth(root,0)
}

func subMaxDepth(node *TreeNode,depth int) int{
	if node==nil {
		return depth
	}
	left:=subMaxDepth(node.Left,depth+1)
	right:=subMaxDepth(node.Right,depth+1)
	if left>right {
		return left
	}else{
		return right
	}
}

// 111. Minimum Depth of Binary Tree
func minDepth(root *TreeNode) int {
	return subMinDepth(root,0)
}

func subMinDepth(node *TreeNode,depth int)  int{
	if node==nil {
		return depth
	}
    if node.Left==nil && node.Right==nil{
        return depth+1
    }
	var left,right int
	if node.Left!=nil {
		left=subMinDepth(node.Left,depth+1)
	}else{
		left=100000
	}
	if node.Right!=nil {
		right=subMinDepth(node.Right,depth+1)
	}else{
		right=100000
	}
	if left<right {
		return left
	}else{
		return right
	}
}

