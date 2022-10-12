/*
 * @lc app=leetcode.cn id=53 lang=cpp
 *
 * [53] 最大子数组和
 */

// @lc code=start
class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int ma=nums[0];
        int lf=0,lz=0;
        for(int i=1;i<nums.size();i++){
            int x=nums[i];
            if(x>=0){
                lz=lz+x;
                if(lz>=ma&&lz>=ma+lf+lz){
                    ma=lz;
                    lf=0;lz=0;
                }else if(ma+lf+lz>=ma&&ma+lf+lz>=lz){
                    ma=ma+lf+lz;
                    lf=0;lz=0;
                }
            }else{
                if(x>ma){
                    ma=x;
                    lf=0;lz=0;
                }else if(lz+x>=0){
                    lz=lz+x;
                }else{
                    lf=lf+lz+x;
                    lz=0;
                }
            }
        }
        return ma;
    }
};
// @lc code=end

