/*
 * @lc app=leetcode.cn id=347 lang=cpp
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
bool cmp(pair<int,int> a,pair<int,int> b){
    return a.second>b.second;
}

class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        map<int,int> m;
        for(int i=0;i<nums.size();i++){
            m[nums[i]]++;
        }
        vector<pair<int,int>> vp(m.begin(),m.end());
        sort(vp.begin(),vp.end(),cmp);
        vector<int> ans;
        for(int i=0;i<k;i++){
            ans.push_back(vp[i].first);
        }
        return ans;
    }
};
// @lc code=end

