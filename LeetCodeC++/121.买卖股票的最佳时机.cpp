/*
 * @lc app=leetcode.cn id=121 lang=cpp
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int ma=0,mi=prices[0];
        for(int i=1;i<prices.size();i++){
            if(prices[i]<=mi){
                mi=prices[i];
            }else{
                ma=max(ma,prices[i]-mi);
            }
        }
        return ma;
    }
};
// @lc code=end

