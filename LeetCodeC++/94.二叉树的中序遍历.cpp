/*
 * @lc app=leetcode.cn id=94 lang=cpp
 *
 * [94] 二叉树的中序遍历
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
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> ret={};
        MidOrderTra(root,ret);
        return ret;
    }

    void MidOrderTra(TreeNode* x,vector<int> &ret){
        if(x==nullptr) return;
        MidOrderTra(x->left,ret);
        ret.push_back(x->val);
        MidOrderTra(x->right,ret);
    }
};
// @lc code=end

