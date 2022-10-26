/*
 * @lc app=leetcode.cn id=98 lang=cpp
 *
 * [98] 验证二叉搜索树
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
    bool isValidBST(TreeNode* root) {
        return !ttt(root,LONG_MIN,LONG_MAX);
    }

    bool ttt(TreeNode* x,long l,long r){
        if(x==nullptr)return false;
        if(x->val>=r||x->val<=l)return true;
        return ttt(x->left,l,x->val)|ttt(x->right,x->val,r);
    }
};
// @lc code=end

