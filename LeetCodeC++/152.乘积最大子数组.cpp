/*
 * @lc app=leetcode.cn id=152 lang=cpp
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
class Solution {
public:
    int maxProduct(vector<int>& nums) {
        // int maxres=INT_MIN;
        // for(int i=0;i<nums.size();i++){
        //     if(maxres<nums[i])maxres=nums[i];
        //     int tempsum=1;
        //     int j;
        //     for(j=i;j>=0&&nums[j]!=0;j--){
        //         tempsum*=nums[j];
        //         if(tempsum>maxres){
        //             maxres=tempsum;
        //         }
        //     }
        //     if(j>=0&&nums[j]==0&&0>maxres){
        //         maxres=0;
        //     }
        // }
        // return maxres;
        int maxF = nums[0], minF = nums[0], ans = nums[0];
        for (int i = 1; i < nums.size(); ++i) {
            int mx = maxF, mn = minF;
            maxF = max(mx * nums[i], max(nums[i], mn * nums[i]));
            minF = min(mn * nums[i], min(nums[i], mx * nums[i]));
            ans = max(maxF, ans);
        }
        return ans;
    }
};
// @lc code=end

