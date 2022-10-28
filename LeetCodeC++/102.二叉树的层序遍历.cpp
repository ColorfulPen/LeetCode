/*
 * @lc app=leetcode.cn id=102 lang=cpp
 *
 * [102] 二叉树的层序遍历
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
    
    vector<vector<int>> levelOrder(TreeNode* root) {
        queue<TreeNode*>q;
        vector<vector<int>> ret={};
        if(root)q.push(root);else return ret;
        vector<int> temp={};
        int num=1;
        while(num){
            TreeNode *x=q.front();
            temp.push_back(x->val);
            if(x->left)q.push(x->left);
            if(x->right)q.push(x->right);
            q.pop();
            num--;
            if(!num){
                ret.push_back(temp);
                num=q.size();
                temp={};
            }
        }
        return ret;

    }
};
// @lc code=end

