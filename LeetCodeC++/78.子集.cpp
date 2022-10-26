/*
 * @lc app=leetcode.cn id=78 lang=cpp
 *
 * [78] 子集
 */

// @lc code=start
class Solution {
public:
    vector<vector<int>> ret={};
    vector<int>got={};
    vector<vector<int>> subsets(vector<int>& nums) {
        recall(nums,0);
        return ret;
    }
    void recall(vector<int>& nums,int index){
        ret.push_back(got);
        for(int i=index;i<nums.size();i++){
            got.push_back(nums[i]);
            recall(nums,i+1);
            got.pop_back();
        }

    }
};
// @lc code=end

