/*
 * @lc app=leetcode.cn id=42 lang=cpp
 *
 * [42] 接雨水
 */

// @lc code=start
class Solution {
public:
    int trap(vector<int>& height) {
        int hm=0;
        stack<int> s;
        int sum=0;
        for(int i=0;i<height.size();i++){
            if(height[i]>=hm){
                while(!s.empty()){
                    sum+=hm-s.top();
                    s.pop();
                }
                hm=height[i];
            }else{
                s.push(height[i]);
            }
        }
        int hsm=0;
        while(!s.empty()){
            if(hsm<s.top()){
                hsm=s.top();
            }else{
                sum+=hsm-s.top();
            }
            s.pop();
        }
        return sum;
    }
};
// @lc code=end

