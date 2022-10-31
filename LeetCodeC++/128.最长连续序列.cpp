/*
 * @lc app=leetcode.cn id=128 lang=cpp
 *
 * [128] 最长连续序列
 */

// @lc code=start
class Solution {
public:
    unordered_map<int,int> m,l;
    int longestConsecutive(vector<int>& nums) {
        int longest=0;
        for(auto x:nums){
            if(m.find(x)==m.end()){
                m[x]=x;l[x]=1;
                if(m.find(x-1)!=m.end()){
                    m[x]=find(x-1);
                    l[m[x]]++;
                }
                if(m.find(x+1)!=m.end()){
                    m[x+1]=find(x);
                    l[m[x]]+=l[x+1];
                }
            }
            longest=longest>l[m[x]]?longest:l[m[x]];
        }
        return longest;
    }
    int find(int i){
        return m[i]==i?i:(m[i]=find(m[i]));
    }
};
// @lc code=end

