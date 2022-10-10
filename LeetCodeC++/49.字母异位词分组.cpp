/*
 * @lc app=leetcode.cn id=49 lang=cpp
 *
 * [49] 字母异位词分组
 */

// @lc code=start
bool mycmp(string a,string b){
    if(a.size()!=b.size())return a.size()<b.size();
    return a<b;
}
class Solution {
public:
    
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        map<string,vector<string>> m;
        vector<vector<string>> v;
        for(int i=0;i<strs.size();i++){
            string k=strs[i];
            sort(k.begin(),k.end());
            m[k].push_back(strs[i]);
        }
        for(auto it=m.begin();it!=m.end();it++){
            v.push_back(it->second);
        }
        return v;
    }
};
// @lc code=end

