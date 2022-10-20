/*
 * @lc app=leetcode.cn id=79 lang=cpp
 *
 * [79] 单词搜索
 */

// @lc code=start
class Solution {
public:
    string w="";
    bool exist(vector<vector<char>>& board, string word) {
        w=word;
        for(int i=0;i<board.size();i++){
            for(int j=0;j<board[0].size();j++){
                if(backtrack(board,0,i,j))return true;
            }
        }
        return false;
    }
    bool backtrack(vector<vector<char>>& board,int index,int h,int l){
        if(index==w.size()-1&&board[h][l]==w[index])return true;
        if(board[h][l]==w[index]){
            bool a=false,b=false,c=false,d=false;
            char temp=board[h][l];
            board[h][l]='!';
            if(h-1>=0){
                a=backtrack(board,index+1,h-1,l);
            }
            if(l-1>=0){
                b=backtrack(board,index+1,h,l-1);
            }
            if(h+1<board.size()){
                c=backtrack(board,index+1,h+1,l);
            }
            if(l+1<board[0].size()){
                d=backtrack(board,index+1,h,l+1);
            }
            board[h][l]=temp;
            return a|b|c|d;
        }
        return false;
    }

};
// @lc code=end

