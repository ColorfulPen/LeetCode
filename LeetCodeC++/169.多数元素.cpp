/*
 * @lc app=leetcode.cn id=169 lang=cpp
 *
 * [169] 多数元素
 */

// @lc code=start
class Solution {
public:
    int majorityElement(vector<int>& nums) {
        int ret=nums[0],num=0;
        for(auto x:nums){
            if(x==ret){
                num++;
            }else if(num==0){
                ret=x;
                num=1;
            }else{
                num--;
            }
        }
        return ret;
    }
};
// @lc code=end

