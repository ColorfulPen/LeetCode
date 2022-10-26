/*
 * @lc app=leetcode.cn id=75 lang=cpp
 *
 * [75] 颜色分类
 */

// @lc code=start
class Solution {
public:
    void sortColors(vector<int>& nums) {
        int r=0,w=0,b=0;
        for(int i=0;i<nums.size();i++){
            if(nums[i]==0)r++;
            else if(nums[i]==1)w++;
            else b++;
        }
        int i=0;
        while(r){
            nums[i]=0;
            i++;
            r--;
        }
        while(w){
            nums[i]=1;
            i++;
            w--;
        }
        while(b){
            nums[i]=2;
            i++;
            b--;
        }
        //正规做法是双指针
    }
};
// @lc code=end

