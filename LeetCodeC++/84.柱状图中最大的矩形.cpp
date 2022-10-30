/*
 * @lc app=leetcode.cn id=84 lang=cpp
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
class Solution {
public:
    int largestRectangleArea(vector<int>& heights) {
        int len=heights.size();
        if(len==0)return 0;
        stack<int> s;
        s.push(0);
        int maxsq=heights[0];
        for(int i=1;i<len;i++){
            if(heights[i]>=heights[s.top()]){
                s.push(i);
            }else{
                while(!s.empty()&&heights[i]<heights[s.top()]){
                    int h=heights[s.top()];
                    s.pop();
                    while(s.size()>0&&h==heights[s.top()]){
                        s.pop();
                    }
                    int sq=h*(i-1-(s.size()!=0?s.top():-1));
                    maxsq=max(maxsq,sq);
                }
                s.push(i);
            }
        }
        int temp=s.top();
        while(s.size()>1){
            int h=heights[s.top()];
            s.pop();
            while(s.size()>1&&h==heights[s.top()]){
                s.pop();
            }
            int sq=h*(temp-s.top());
            maxsq=max(maxsq,sq);
        }
        int w=len*heights[s.top()];
        maxsq=max(maxsq,w);
        return maxsq;
    }
};
// @lc code=end

