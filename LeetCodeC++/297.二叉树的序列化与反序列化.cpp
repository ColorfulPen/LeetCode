/*
 * @lc app=leetcode.cn id=297 lang=cpp
 *
 * [297] 二叉树的序列化与反序列化
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
class Codec {
public:

    // Encodes a tree to a single string.
    string serialize(TreeNode* root) {
        queue<TreeNode*> q;
        string ret="";
        q.push(root);
        while(!q.empty()){
            TreeNode*x=q.front();
            q.pop();
            ret+=x?to_string(x->val)+",":"null,";
            if(x){
                q.push(x->left);
                q.push(x->right);
            }
        }
        return ret;
    }

    // Decodes your encoded data to tree.
    TreeNode* deserialize(string data) {
        
        vector<string> vs=split(data,",");
        if(vs[0]=="null") return nullptr;
        queue<TreeNode*> q;
        TreeNode* ret=new TreeNode(stoi(vs[0]));
        q.push(ret);
        int i=1;
        while(i<vs.size()){
            TreeNode* x=q.front();
            q.pop();
            if(vs[i]!="null"){
                TreeNode* temp=new TreeNode(stoi(vs[i]));
                x->left=temp;
                q.push(temp);
            }
            i++;
            if(vs[i]!="null"){
                TreeNode* temp=new TreeNode(stoi(vs[i]));
                x->right=temp;
                q.push(temp);
            }
            i++;
        }
        return ret;
    }
    vector<string> split(const string& str, const string& delim) {
        vector<string> res;
        if("" == str) return res;
        //先将要切割的字符串从string类型转换为char*类型
        char * strs = new char[str.length() + 1] ; //不要忘了
        strcpy(strs, str.c_str());
        char * d = new char[delim.length() + 1];
        strcpy(d, delim.c_str());
        char *p = strtok(strs, d);
        while(p) {
            string s = p; //分割得到的字符串转换为string类型
            res.push_back(s); //存入结果数组
            p = strtok(NULL, d);
        }
        return res;
    }
};

// Your Codec object will be instantiated and called as such:
// Codec ser, deser;
// TreeNode* ans = deser.deserialize(ser.serialize(root));
// @lc code=end

