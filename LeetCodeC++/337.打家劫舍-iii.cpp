/*
 * @lc app=leetcode.cn id=337 lang=cpp
 *
 * [337] 打家劫舍 III
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
    int rob(TreeNode* root) {
        vector<int> x=dfs(root);
        return x[0];
    }

    vector<int> dfs(TreeNode* n){
        if(n==nullptr){
            return {0,0,0};
        }
        vector<int> ret(3);
        vector<int> l=dfs(n->left);
        vector<int> r=dfs(n->right);
        ret[0]=max(l[1]+l[2]+r[1]+r[2]+n->val,l[0]+r[0]);
        ret[1]=l[0];
        ret[2]=r[0];
        return ret;
    }
};
// @lc code=end

