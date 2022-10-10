/*
 * @lc app=leetcode.cn id=48 lang=cpp
 *
 * [48] 旋转图像
 */

// @lc code=start
class Solution {
public:
    void rotate(vector<vector<int>>& matrix) {
        int ma=matrix.size()-1;
        int le=matrix.size();
        while(le>1){
            int i=(matrix.size()-le)/2;
            for(int j=i;j<le+i-1;j++){
                int temp1=matrix[j][ma-i];
                matrix[j][ma-i]=matrix[i][j];
                int temp2=matrix[ma-i][ma-j];
                matrix[ma-i][ma-j]=temp1;
                int temp3=matrix[ma-j][i];
                matrix[ma-j][i]=temp2;
                matrix[i][j]=temp3;
            }
            le=le-2;
        }
    }
};
// @lc code=end

