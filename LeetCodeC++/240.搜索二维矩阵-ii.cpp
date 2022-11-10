/*
 * @lc app=leetcode.cn id=240 lang=cpp
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
class Solution {
private:
    int h,l;
    // vector<vector<bool>>visited;
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
    //     h=matrix.size(),l=matrix[0].size();
    //     visited.resize(h,vector<bool>(l,false));
    //     return dfs(matrix,0,0,target);
    // }

    // bool dfs(vector<vector<int>>& matrix,int i,int j,int target){
    //     if(i>=h||j>=l||visited[i][j])return false;
    //     if(matrix[i][j]==target)return true;
    //     if(matrix[i][j]<target){
    //         visited[i][j]=true;
    //         return dfs(matrix,i+1,j,target)||dfs(matrix,i,j+1,target);
    //     }
    //     return false;
        h=0,l=matrix[0].size()-1;
        while(h<matrix.size()&&l>=0){
            if(matrix[h][l]==target)return true;
            else if(matrix[h][l]>target)l--;
            else h++;
        }
        return false;
    }
};
// @lc code=end

