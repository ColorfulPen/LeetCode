/*
 * @lc app=leetcode.cn id=39 lang=cpp
 *
 * [39] 组合总和
 */

// @lc code=start
class Solution {
public:
    vector<vector<int>> ret={};
    vector<int> candidates={};
    int target=0;
    vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
        this->candidates=candidates;
        this->target=target;
        for(int i=0;i<candidates.size();i++){
            backtrack({candidates[i]},candidates[i],i);
        }
        return ret;
    }
    void backtrack(vector<int>trace,int sum,int pos){
        if(sum==target) ret.push_back(trace);
        else if(sum>target)return;
        else{
            for(int i=pos;i<candidates.size();i++){
                trace.push_back(candidates[i]);
                backtrack(trace,sum+candidates[i],i);
                trace.pop_back();
            }
        }
    }
};
// @lc code=end

