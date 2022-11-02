/*
 * @lc app=leetcode.cn id=124 lang=cpp
 *
 * [124] 二叉树中的最大路径和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    int maxsum=INT_MIN;
    int maxPathSum(TreeNode* root) {
        postorder(root);
        return maxsum;
    }
    int postorder(TreeNode* root){
        if(root==nullptr)return 0;
        int leftsum=postorder(root->left);
        int rightsum=postorder(root->right);
        maxsum=max(maxsum,(leftsum>0?leftsum:0)+(rightsum>0?rightsum:0)+root->val);
        int branchmax=max(leftsum,rightsum);
        return root->val+(branchmax>0?branchmax:0);
    }
};
// @lc code=end

