/*
 * @lc app=leetcode.cn id=207 lang=cpp
 *
 * [207] 课程表
 */

// @lc code=start
class Solution {

public:
    unordered_map<int,vector<int>> graph;
    unordered_map<int,int>indegree;
    unordered_map<int,bool>visited;
    bool canFinish(int numCourses, vector<vector<int>>& prerequisites) {
        for(auto x:prerequisites){
            indegree[x[0]]++;
            visited[x[0]]=false;
            graph[x[0]].push_back(x[1]);
        }
        for(auto x:graph){
            if(dfs(x.first))return false;
        }
        return true;
    }
    bool dfs(int x){
        bool ret=false;
        if(indegree.find(x)==indegree.end()||indegree[x]==0)return false;
        else if(visited[x]) return true;
        else{
            visited[x]=true;
            for(auto temp:graph[x]){
                ret=ret|dfs(temp);
                indegree[x]--;
            }
        }
        return ret;
    }
};
// @lc code=end

