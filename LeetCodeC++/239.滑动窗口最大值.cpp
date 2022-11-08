/*
 * @lc app=leetcode.cn id=239 lang=cpp
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
class Solution {
public:
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        deque<int> deq;
        int len=nums.size();
        vector<int> ret;
        for(int i=0;i<k;i++){
            while(!deq.empty()&&nums[deq.back()]<=nums[i]){
                deq.pop_back();
            }
            deq.push_back(i);
        }
        ret.push_back(nums[deq.front()]);
        for(int i=k;i<len;i++){
            while(!deq.empty()&&nums[deq.back()]<=nums[i]){
                deq.pop_back();
            }
            deq.push_back(i);
            while(deq.front()<=i-k){
                deq.pop_front();
            }
            
            ret.push_back(nums[deq.front()]);
        }
        return ret;
    }
};
// @lc code=end

