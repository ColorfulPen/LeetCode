/*
 * @lc app=leetcode.cn id=704 lang=cpp
 *
 * [704] 二分查找
 */

// @lc code=start
class Solution {
public:
    int search(vector<int>& nums, int target) {
        int l=0,r=nums.size()-1;
        int m;
        while(l<=r){
            m=(l+r)>>1;
            if(nums[m]==target)return m;
            else if(nums[m]>target)r=m-1;
            else l=m+1;
        }
        return -1;
    }
};
// @lc code=end

