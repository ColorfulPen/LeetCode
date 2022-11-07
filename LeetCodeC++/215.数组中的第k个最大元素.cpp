/*
 * @lc app=leetcode.cn id=215 lang=cpp
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        priority_queue<int,vector<int>,greater<int>>small_heap;
        int num=0;
        for(auto x:nums){
            small_heap.push(x);
            num++;
            if(num>k)small_heap.pop();
        }
        return small_heap.top();
    }
};
// @lc code=end

