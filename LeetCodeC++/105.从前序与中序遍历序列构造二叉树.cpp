/*
 * @lc app=leetcode.cn id=105 lang=cpp
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        int pi=0;
        TreeNode* x=buildB(preorder,inorder,pi,preorder.size());
        return x;
    }
    TreeNode* buildB(vector<int>& preorder, vector<int>& inorder,int &pi,int mid){
        if(pi>=preorder.size())return nullptr;
        int ii=find(inorder.begin(),inorder.end(),preorder[pi])-inorder.begin();
        TreeNode* x;
        if(ii<mid){
            x=new TreeNode(preorder[pi]);
            pi++;
            x->left=buildB(preorder,inorder,pi,ii);
            x->right=buildB(preorder,inorder,pi,mid);
        }else{
            x=nullptr;
        }
        return x;
    }
};
// @lc code=end

