/*
 * @lc app=leetcode.cn id=399 lang=cpp
 *
 * [399] 除法求值
 */

// @lc code=start
class Solution {
public:
    bool dfs(unordered_map<string,vector<pair<string,double>>> mp,string a, string b, double &sum,unordered_map<string,int> flag){   
        // cout<<"a="<<a<<" b="<<b<<" sum="<<sum<<endl;
        if (a==b) {
            return true;
        }
        for (auto [p0,p1]:mp[a]){
            // cout<<"p0="<<p0<<" p1="<<p1<<endl;
                if (flag[p0]==0){
                    flag[p0]=1;
                    sum *= p1;
                    if (dfs(mp,p0,b,sum,flag)){
                        return true;
                    }
                    sum /= p1;
                    flag[p0]=0;
                }
        }
        return false;
    }
    vector<double> calcEquation(vector<vector<string>>& equations, vector<double>& values, vector<vector<string>>& queries) {
        unordered_map<string,vector<pair<string,double>>> mp;
        int n = equations.size();
        for (int i =0; i<n; i++){
            mp[equations[i][0]].emplace_back(pair(equations[i][1],values[i]));
            mp[equations[i][1]].emplace_back(pair(equations[i][0],1.0/values[i]));           
        }
        vector<double> ans(queries.size());
        int idx = 0;
        for (auto querie : queries){
            if (mp.find(querie[0])!=mp.end() && mp.find(querie[1])!=mp.end()){
                    unordered_map<string,int> flag;
                    double sum =1.0;
                    ans[idx++] = dfs(mp,querie[0],querie[1],sum,flag)==true ? sum:-1.0 ;
            }
            else{
                ans[idx++] = -1.0;
            }
        }
        return ans;
    }
};
//从一点出发深度搜索
// @lc code=end

