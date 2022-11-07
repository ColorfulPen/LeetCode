/*
 * @lc app=leetcode.cn id=236 lang=cpp
 *
 * [236] 二叉树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    TreeNode* ret=nullptr;
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if(root==nullptr)return nullptr;
        dfs(root,p,q);
        return ret;
    }
    int dfs(TreeNode* root, TreeNode* p, TreeNode* q){
        if(root==nullptr)return 0;
        if(ret!=nullptr)return 2;
        int find=dfs(root->left,p,q)+dfs(root->right,p,q);
        if(ret!=nullptr)return 2;
        if(root==p||root==q)
            find++;
        if(find==2)ret=root;
        return find;
    }
};
// @lc code=end

