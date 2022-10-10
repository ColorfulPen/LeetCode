/*
 * @lc app=leetcode.cn id=9 lang=cpp
 *
 * [9] 回文数
 */

// @lc code=start
class Solution {
public:
    bool isPalindrome(int x) {
        if(x<0)return false;
        long a=0;
        int b=x;
        while(b){
            a=a*10+b%10;
            b/=10;
        }
        if(a==x)return true;
        return false;
    }
};
// @lc code=end

