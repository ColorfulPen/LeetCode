/*
 * @lc app=leetcode.cn id=238 lang=cpp
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
class Solution {
public:
    vector<int> productExceptSelf(vector<int>& nums) {
        int len=nums.size();
        vector<int> ans(len);
        int lm=nums[0];
        ans[len-1]=1;
        for(int i=len-2;i>=0;i--){
            ans[i]=ans[i+1]*nums[i+1];
        }
        for(int i=1;i<len;i++){
            ans[i]=lm*ans[i];
            lm=lm*nums[i];
        }
        return ans;
    }
};
// @lc code=end

