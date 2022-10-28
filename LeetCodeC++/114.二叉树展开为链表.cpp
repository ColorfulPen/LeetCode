/*
 * @lc app=leetcode.cn id=114 lang=cpp
 *
 * [114] 二叉树展开为链表
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
    void flatten(TreeNode* root) {
        d(root);
    }
    TreeNode* d(TreeNode* x){
        if(x==nullptr)return x;
        TreeNode* temp=x->right;
        if(x->left!=nullptr){
            x->right=x->left;
            x->left=nullptr;
            x=x->right;
            x=d(x);
        }
        if(temp!=nullptr){
            x->right=temp;
            x=x->right;
            x=d(x);
        }
        return x;
    }

};
// @lc code=end

