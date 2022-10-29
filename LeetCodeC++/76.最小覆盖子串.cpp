/*
 * @lc app=leetcode.cn id=76 lang=cpp
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
class Solution {
public:
    // unordered_map <char, int> ori, cnt;
    // bool check() {
    //     for (const auto &p: ori) {
    //         if (cnt[p.first] < p.second) {
    //             return false;
    //         }
    //     }
    //     return true;
    // }
    // string minWindow(string s, string t) {
    //     for (const auto &c: t) {
    //         ++ori[c];
    //     }
    //     int l = 0, r = -1;
    //     int len = INT_MAX, ansL = -1, ansR = -1;
    //     while (r < int(s.size())) {
    //         if (ori.find(s[++r]) != ori.end()) {
    //             ++cnt[s[r]];
    //         }
    //         while (check() && l <= r) {
    //             if (r - l + 1 < len) {
    //                 len = r - l + 1;
    //                 ansL = l;
    //             }
    //             if (ori.find(s[l]) != ori.end()) {
    //                 --cnt[s[l]];
    //             }
    //             ++l;
    //         }
    //     }

    //     return ansL == -1 ? string() : s.substr(ansL, len);
    // }

    unordered_map<char,int> sm,tm;
    bool check(){
        for(const auto &p : tm){
            if(sm[p.first]<p.second)return false;
        }
        return true;
    }

    string minWindow(string s, string t) {
        for(const auto &c:t){
            ++tm[c];
        }
        int l=0,r=-1;
        int ans=-1;
        int minl=INT_MAX;
        while(r<int(s.size())){
            //size()函数返回size_t是无符号的
            if(tm.find(s[++r])!=tm.end()) ++sm[s[r]];
            while(check()&&l<=r){
                if(r-l+1<minl){
                    minl=r-l+1;
                    ans=l;
                }
                
                if(tm.find(s[l])!=tm.end())
                    --sm[s[l]];
                ++l;
            }
        }
        return ans==-1?string():s.substr(ans,minl);
    }
};
// @lc code=end

