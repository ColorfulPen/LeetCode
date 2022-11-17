/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func depth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    if root.Left == nil && root.Right == nil {
        return 1
    }

    var dp int
    if root.Left != nil {
        dp = depth(root.Left)
    }

    if root.Right != nil && depth(root.Right)>dp {
        dp = depth(root.Right)
    }

    return dp+1
}
func countNodes(root *TreeNode) int {

    if root == nil {
        return 0
    }

    num := 1
   
    for root != nil {
        if root.Left==nil && root.Right == nil {
            break
        }
        if root.Left != nil {
            if root.Right != nil {
                ldp := depth(root.Left)
                rdp := depth(root.Right)
                if ldp == rdp {
                   root = root.Right
                   num = num * 2 + 1
                   continue
                }
                root = root.Left
                num = num *2
                continue
            }
            root = root.Left
            num = num * 2
            continue
        }   
    }

    return num
}

// @lc code=end

