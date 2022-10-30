/*
 * @lc app=leetcode.cn id=85 lang=cpp
 *
 * [85] 最大矩形
 */

// @lc code=start
class Solution {
public:
    int maximalRectangle(vector<vector<char>>& matrix) {
        int h=matrix.size();
        int l=matrix[0].size();
        vector<vector<int>>ln(h+2,vector<int>(l,0));
        for(int i=1;i<h+1;i++){
            for(int j=0;j<l;j++){
                if(matrix[i-1][j]=='1'){
                    if(j==0){
                        ln[i][j]=1;
                    }else{
                        ln[i][j]=ln[i][j-1]+1;
                    }
                }   
            }
        }
        int ans=0;
        for(int i=0;i<l;i++){
            stack<int> s;
            s.push(0);
            for(int j=1;j<h+2;j++){
                while(ln[j][i]<ln[s.top()][i]){
                    int temp=s.top();
                    s.pop();
                    ans=max(ans,(j-1-s.top())*ln[temp][i]);
                }
                s.push(j);
            }
        }
        return ans;
    }
};
// @lc code=end

