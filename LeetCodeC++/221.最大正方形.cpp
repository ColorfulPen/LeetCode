/*
 * @lc app=leetcode.cn id=221 lang=cpp
 *
 * [221] 最大正方形
 */

// @lc code=start
class Solution {
public:
    int maximalSquare(vector<vector<char>>& matrix) {
    //     int maxsize=0;
    //     int h=matrix.size();
    //     int l=matrix[0].size();
    //     vector<vector<int>>dp(h,vector<int>(l,0));
    //     for(int i=0;i<h;i++){
    //         dp[i][0]=matrix[i][0]-'0';
    //         if(dp[i][0])maxsize=1;
    //     }
    //     for(int i=0;i<l;i++){
    //         dp[0][i]=matrix[0][i]-'0';
    //         if(dp[0][i])maxsize=1;
    //     }
    //     for(int i=1;i<h;i++){
    //         for(int j=1;j<l;j++){
    //             dp[i][j]=panduan(matrix,i,j,dp[i-1][j-1]);
    //             maxsize=max(maxsize,dp[i][j]);
    //         }
    //     }
    //     return maxsize*maxsize;
    // }
    // int panduan(vector<vector<char>>& matrix,int i,int j,int len){
    //     if(matrix[i][j]-'0'==0)return 0;
    //     for(int k=0;k<=len;k++){
    //         if(!((matrix[i][j-k]-'0')&&(matrix[i-k][j]-'0'))){
    //             return k;
    //         }
    //     }
    //     return len+1;

        // int maxsize=0;
        // int h=matrix.size();
        // int l=matrix[0].size();
        // vector<vector<int>>dp(h,vector<int>(l,0));
        // for(int i=0;i<h;i++){
        //     for(int j=0;j<l;j++){
        //         if(matrix[i][j]=='1'){
        //             if(i==0||j==0)
        //                 dp[i][j]=1;
        //             else
        //                 dp[i][j]=min(dp[i-1][j],min(dp[i-1][j-1],dp[i][j-1]))+1;
        //             maxsize=max(maxsize,dp[i][j]);
        //         }
        //     }
        // }
        // return maxsize*maxsize;



        if (matrix.size() == 0 || matrix[0].size() == 0) {
            return 0;
        }
        int maxSide = 0;
        int rows = matrix.size(), columns = matrix[0].size();
        vector<vector<int>> dp(rows, vector<int>(columns));
        for (int i = 0; i < rows; i++) {
            for (int j = 0; j < columns; j++) {
                if (matrix[i][j] == '1') {
                    if (i == 0 || j == 0) {
                        dp[i][j] = 1;
                    } else {
                        dp[i][j] = min(min(dp[i - 1][j], dp[i][j - 1]), dp[i - 1][j - 1]) + 1;
                    }
                    maxSide = max(maxSide, dp[i][j]);
                }
            }
        }
        int maxSquare = maxSide * maxSide;
        return maxSquare;

    }
};
// @lc code=end

