/*
 * @lc app=leetcode.cn id=56 lang=cpp
 *
 * [56] 合并区间
 */

// @lc code=start
bool mycmp(vector<int>a,vector<int>b){
    if(a[0]==b[0])return a[1]<b[1];
    return a[0]<b[0];
}
class Solution {
public:
    vector<vector<int>> merge(vector<vector<int>>& intervals) {
        // sort(intervals.begin(),intervals.end(),mycmp);
        // for(int i=1;i<intervals.size();i++){
        //     if(intervals[i][0]<=intervals[i-1][1]){
        //         if(intervals[i][1]>intervals[i-1][1]){
        //             intervals[i-1][1]=intervals[i][1];
        //         }
        //         intervals.erase(intervals.begin()+i);
        //         i--;
        //     }
        // }
        // return intervals;
        int n = intervals.size();
        vector<vector<int>> res;
        vector<int> starts, ends;
        for (int i = 0; i < n; ++i) {
            starts.push_back(intervals[i][0]);
            ends.push_back(intervals[i][1]);
        }
        sort(starts.begin(), starts.end());
        sort(ends.begin(), ends.end());
        for (int i = 0, j = 0; i < n; ++i) {
            if (i == n - 1 || starts[i + 1] > ends[i]) {
                res.push_back({starts[j], ends[i]});
                j=i+1;
            }
            //真牛哇
        } 
        return res;
    }
};
// @lc code=end

