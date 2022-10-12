/*
 * @lc app=leetcode.cn id=1790 lang=cpp
 *
 * [1790] 仅执行一次字符串交换能否使两个字符串相等
 */

// @lc code=start
class Solution {
public:
    bool areAlmostEqual(string s1, string s2) {
        char c1,c2,c3,c4;
        int dnum=0;
        for(int i=0;i<s1.size();i++){
            if(s1[i]!=s2[i]){
                if(dnum==0){
                    c1=s1[i];
                    c3=s2[i];
                    dnum++;
                }else if(dnum==1){
                    c2=s1[i];
                    c4=s2[i];
                    dnum++;
                }else{
                    dnum++;
                }
            }
        }
        if(dnum==0||(c1==c4&&c2==c3&&dnum==2))return true;
        return false;
    }
};
// @lc code=end

