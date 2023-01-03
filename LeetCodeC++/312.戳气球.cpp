/*
 * @lc app=leetcode.cn id=312 lang=cpp
 *
 * [312] 戳气球
 */

// @lc code=start
class Solution {
public:
    int maxCoins(vector<int>& nums) {
        int n=nums.size();
        vector<vector<int>> dp(n+2,vector<int>(n+2,0));
        vector<int>newnums(n+2);
        newnums[0]=1;
        newnums[n+1]=1;
        for(int i=n-1;i>=0;i--){
            newnums[i+1]=nums[i];
        }
        for(int k=2;k<n+2;k++){
            for(int i=0;i+k<n+2;i++){
                for(int j=i+1;j<i+k;j++){
                    dp[i][i+k]=max(dp[i][i+k],newnums[i]*newnums[j]*newnums[i+k]+dp[i][j]+dp[j][i+k]);
                }
            }
        }
        return dp[0][n+1];
    }
};
// @lc code=end

