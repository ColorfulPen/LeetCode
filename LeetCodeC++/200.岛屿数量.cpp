/*
 * @lc app=leetcode.cn id=200 lang=cpp
 *
 * [200] 岛屿数量
 */

// @lc code=start
class Solution {
public:
    
    int numIslands(vector<vector<char>>& grid) {
        vector<vector<bool>> tag(grid.size(),vector<bool>(grid[0].size(),false));
        int m=grid.size();
        int n=grid[0].size();
        int ret=0;
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                if(grid[i][j]=='1'&&!tag[i][j]){
                    deep(grid,tag,i,j);
                    ret++;
                }
            }
        }
        return ret;
    }
    void deep(vector<vector<char>>& grid,vector<vector<bool>> &tag,int i,int j){
        if(i<0||i>=grid.size()||j<0||j>=grid[0].size())return;
        if(grid[i][j]=='1'&&!tag[i][j]){
            tag[i][j]=true;
            deep(grid,tag,i-1,j);
            deep(grid,tag,i+1,j);
            deep(grid,tag,i,j-1);
            deep(grid,tag,i,j+1);
        }
    }

};
// @lc code=end

