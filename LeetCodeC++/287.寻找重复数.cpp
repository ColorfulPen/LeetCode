/*
 * @lc app=leetcode.cn id=287 lang=cpp
 *
 * [287] 寻找重复数
 */

// @lc code=start
class Solution {
public:
    int findDuplicate(vector<int>& nums) {
        // int n=nums.size();
        // int l=1,r=n-1;
        // int ans=nums[0];
        // while(l<=r){
        //     int mid=(l+r)>>1;
        //     int cnt=0;
        //     for(int i=0;i<n;i++){
        //         cnt+=nums[i]<=mid;
        //     }
        //     if(cnt>mid){
        //         r=mid-1;
        //         ans=mid;
        //     }else{
        //         l=mid+1;
        //     }   
        // }
        // return ans;

        int f=nums[0],s=nums[0];
        do{
            f=nums[nums[f]];
            s=nums[s];
        }while(f!=s);
        f=nums[0];
        while(f!=s){
            f=nums[f];
            s=nums[s];
        }
        return s;
    }
};
// @lc code=end

