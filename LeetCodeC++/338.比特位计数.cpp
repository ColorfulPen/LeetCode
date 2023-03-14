/*
 * @lc app=leetcode.cn id=338 lang=cpp
 *
 * [338] 比特位计数
 */

// @lc code=start
class Solution {
public:
    vector<int> countBits(int n) {
        vector<int> ans(n+1);
        ans[0]=0;
        int temp1=1;
        int temp2=1;
        for(int i=1;i<=n;i++){
            if(i==temp1){
                temp2=temp1;
                temp1*=2;
            }
            ans[i]=ans[i-temp2]+1;
        }
        return ans;
    }
};
// @lc code=end

