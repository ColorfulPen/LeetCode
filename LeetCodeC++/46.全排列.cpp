/*
 * @lc app=leetcode.cn id=46 lang=cpp
 *
 * [46] 全排列
 */

// @lc code=start
class Solution {
public:
    vector<vector<int>> ret={};
    vector<vector<int>> permute(vector<int>& nums) {
        recursion(nums,{});
        return ret;
    }
    void recursion(vector<int> rest,vector<int> done){
        if(rest.size()==0){
            ret.push_back(done);
        }
        for(int i=0;i<rest.size();i++){
            int x=rest[i];
            done.push_back(x);
            rest.erase(rest.begin()+i);
            recursion(rest,done);
            rest.insert(rest.begin()+i,x);
            done.pop_back();
        }
    }

};
// @lc code=end

