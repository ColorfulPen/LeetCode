/*
 * @lc app=leetcode.cn id=300 lang=cpp
 *
 * [300] 最长递增子序列
 */

// @lc code=start
class Solution {
public:
    int lengthOfLIS(vector<int>& nums) {
        int n=nums.size();
        int len=1;
        vector<int> d(n);
        d[0]=nums[0];
        for(int i=1;i<n;i++){
            if(nums[i]>d[len-1]){
                d[len]=nums[i];
                len++;
            }else if(nums[i]<d[len-1]){
                int l=0,r=len-1;
                int mid;
                while(l<=r){
                    mid=(l+r)>>1;
                    if(d[mid]>=nums[i]){
                        r=mid-1;
                    }else{
                        l=mid+1;
                    }
                }
                d[r+1]=nums[i];
            }
        }
        return len;
    }
};
// @lc code=end

