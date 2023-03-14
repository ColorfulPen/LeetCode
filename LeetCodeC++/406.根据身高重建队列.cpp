/*
 * @lc app=leetcode.cn id=406 lang=cpp
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
class Solution {
public:
    vector<vector<int>> reconstructQueue(vector<vector<int>>& people) {
        sort(people.begin(),people.end(),[](const vector<int> &a,const vector<int> b)->bool {return a[0]>b[0]||(a[0]==b[0]&&a[1]<b[1]);});
        vector<vector<int>> ans;
        for(auto person:people){
            ans.insert(ans.begin()+person[1],person);
        }
        return ans;
    }
};
// @lc code=end

