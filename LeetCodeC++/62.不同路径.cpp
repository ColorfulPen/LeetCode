/*
 * @lc app=leetcode.cn id=62 lang=cpp
 *
 * [62] 不同路径
 */

// @lc code=start
class Solution {
public:
    int uniquePaths(int m, int n) {
        vector<vector<int>> x(m+1,vector<int>(n+1,0));
        x[0][1]=1;
        for(int i=1;i<=m;i++){
            for(int j=1;j<=n;j++){
                x[i][j]=x[i-1][j]+x[i][j-1];
            }
        }
        return x[m][n];
    }
};
// @lc code=end

