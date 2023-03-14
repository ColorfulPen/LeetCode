/*
 * @lc app=leetcode.cn id=394 lang=cpp
 *
 * [394] 字符串解码
 */

// @lc code=start
class Solution {
public:
    string getpart(string s,int i,int &reti){
        string rets;
        while(i<s.size()&&s[i]!=']'){
            if(s[i]<='z'&&s[i]>='a'){
                rets+=s[i];
                i++;
            }else if(s[i]>='1'&&s[i]<='9'){
                int time=s[i]-'0';
                i++;
                while(s[i]>='0'&&s[i]<='9'){
                    time=time*10+s[i]-'0';
                    i++;
                }
                if(s[i]=='['){
                    i++;
                    string temps=getpart(s,i,i);
                    for(int j=0;j<time;j++){
                        rets+=temps;
                    }
                    i++;
                }
            }   
        }
        reti=i;
        return rets;
    }
    string decodeString(string s) {
        string ans;
        int i=0;
        return getpart(s,i,i);
    }
};
// @lc code=end

