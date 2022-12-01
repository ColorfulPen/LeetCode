/*
 * @lc app=leetcode.cn id=301 lang=cpp
 *
 * [301] 删除无效的括号
 */

// @lc code=start
class Solution {
// private:
//     vector<unordered_set<string>> vvs;
public:
    // vector<string> removeInvalidParentheses(string s) {
    //     int len=s.size();
    //     vvs.resize(len+1);
    //     reverse(s,0,0,"");
    //     for(int i=len;i>=0;i--){
    //         if(!vvs[i].empty()){
    //             vector<string> ans;
    //             for(auto x:vvs[i]){
    //                 ans.push_back(x);
    //             }
    //             return ans;
    //         }
    //     }
    //     return {""};
    // }
    // void reverse(string s,int index,int num,string temps){
    //     if(num<0)return;
    //     if(index==s.size()){
    //         if(num==0) vvs[temps.size()].insert(temps);
    //         return;
    //     }
    //     reverse(s,index+1,num,temps);//删除
    //     //不删
    //     if(s[index]=='('){
    //         reverse(s,index+1,num+1,temps+s[index]);
    //     }else if(s[index]==')'){
    //         reverse(s,index+1,num-1,temps+s[index]);
    //     }else{
    //         reverse(s,index+1,num,temps+s[index]);
    //     }
    // }
    vector<string> res;
    vector<string> removeInvalidParentheses(string s) {
        int lr=0,rr=0;
        for(auto c:s){
            if(c=='('){
                lr++;
            }else if(c==')'){
                if(lr>0)lr--;
                else rr++;
            }
        }   
        helper(s,0,lr,rr);
        return res;
    } 
    void helper(string s,int index,int lr,int rr){
        if(lr==0&&rr==0){
            if(valid(s))res.push_back(s);
            return;
        }
        for(int i=index;i<s.size();i++){
            if(i!=index&&s[i]==s[i-1]) continue;
            if(lr+rr>s.size()-i) return;
            if(s[i]=='('&&lr>0)helper(s.substr(0,i)+s.substr(i+1),i,lr-1,rr);
            if(s[i]==')'&&rr>0)helper(s.substr(0,i)+s.substr(i+1),i,lr,rr-1);
        }
    }
    
    
    bool valid(string s){
        int num=0;
        for(auto c:s){
            if(c=='(')num++;
            else if(c==')'){
                num--;
                if(num<0)return false;
            }
        }
        return num==0;
    }
};
// @lc code=end

