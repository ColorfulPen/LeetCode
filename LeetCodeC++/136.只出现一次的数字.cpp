/*
 * @lc app=leetcode.cn id=136 lang=cpp
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
class Solution {
public:
    int singleNumber(vector<int>& nums) {
        int x=0;
        for(auto num:nums){
            x^=num;
        }
        return x;
    }
};
// @lc code=end

