/*
 * @lc app=leetcode.cn id=55 lang=cpp
 *
 * [55] 跳跃游戏
 */

// @lc code=start
class Solution {
public:
    bool canJump(vector<int>& nums) {
        // vector<bool> dp(nums.size(),false);
        // dp[0]=true;
        // for(int i=1;i<nums.size();i++){
        //     for(int j=i-1;j>=0;j--){
        //         if(dp[j]&&j+nums[j]>=i){
        //             dp[i]=true;
        //             break;
        //         }
        //     }
        // }
        // if(*dp.rbegin())return true;
        // return false;
        int max=nums[0];
        for(int i=1;i<nums.size();i++){
            if(max<i)return false;
            if(i+nums[i]>max)max=i+nums[i];
        }
        return true;
    }
};
// @lc code=end

