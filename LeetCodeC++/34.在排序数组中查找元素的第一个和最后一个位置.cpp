/*
 * @lc app=leetcode.cn id=34 lang=cpp
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
class Solution {
public:
    vector<int> searchRange(vector<int>& nums, int target) {
        int l=0,r=nums.size()-1;
        int mid=0;
        int m1=-1,m2=-1;
        while(l<=r){
            mid=(l+r)/2;
            if(nums[mid]==target)break;
            else if(nums[mid]<target)l=mid+1;
            else r=mid-1;
        }
        if(l<=r){
            int l1=l,r1=mid;
            int l2=mid,r2=r;
            m1=(l1+r1)/2;
            m2=(l2+r2)/2;
            while(l1<r1){
                m1=(l1+r1)/2;
                if(nums[m1]==target){
                    r1=m1;
                }
                else {
                    if(l1==m1){
                        m1++;
                        break;
                    }
                    l1=m1;
                }
            }
            while(l2<r2){
                m2=(l2+r2+1)/2;
                if(nums[m2]==target){
                    l2=m2;
                }
                else {
                    if(r2==m2){
                        m2--;
                        break;
                    }
                    r2=m2;
                }
            }
        }
        vector<int> x={m1,m2};
        return x;
    }
};
// @lc code=end

