/*
 * @lc app=leetcode.cn id=8 lang=cpp
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
class Solution {
public:
    int myAtoi(string s) {
        long ret=0;
        int i=0;
        int flag=1;
        while(isspace(s[i])){
            i++;
        }
        if(s[i]=='-'){
            flag=-1;
            i++;
        }else if(s[i]=='+'){
            flag=1;
            i++;
        }
        while(isdigit(s[i])){
            ret=ret*10+(s[i]-'0')*flag;
            i++;
            if(ret>INT_MAX)return INT_MAX;
            else if(ret<INT_MIN)return INT_MIN;
        }
        return (int)ret;

    }
};
// @lc code=end

