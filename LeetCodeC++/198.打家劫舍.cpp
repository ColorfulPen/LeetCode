/*
 * @lc app=leetcode.cn id=198 lang=cpp
 *
 * [198] 打家劫舍
 */

// @lc code=start
class Solution {
public:
    int rob(vector<int>& nums) {
        int len=nums.size();
        vector<int> dp(len,0);
        dp[0]=nums[0];
        if(len>=2)dp[1]=max(nums[0],nums[1]);
        if(len>=3)dp[2]=max(nums[0]+nums[2],nums[1]);
        for(int i=3;i<len;i++){
            dp[i]=max(dp[i-1],dp[i-2]+nums[i]);
        }
        return dp[len-1];
    }
};
// @lc code=end

