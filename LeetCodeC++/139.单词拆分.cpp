/*
 * @lc app=leetcode.cn id=139 lang=cpp
 *
 * [139] 单词拆分
 */

// @lc code=start
class Solution {
public:
    bool wordBreak(string s, vector<string>& wordDict) {
        vector<bool> dp(s.size()+1,false);
        dp[0]=true;
        for(int i=1;i<s.size()+1;i++){
            bool flag=false;
            for(int k=1;i-k>=0;k++){
                if(dp[i-k]){
                    string temp=s.substr(i-k,k);
                    for(auto x:wordDict){
                        if(temp==x){
                            flag=true;
                            dp[i]=true;
                            break;
                        }
                    }
                }
                if(flag)break;
            }
        }
        return dp[s.size()];
    }
};
// @lc code=end

