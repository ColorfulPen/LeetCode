/*
 * @lc app=leetcode.cn id=309 lang=cpp
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int n=prices.size();
        vector<int> buy(n,0);
        vector<int> sell(n,0);
        vector<int> stop(n,0);
        buy[0]=-prices[0];
        for(int i=1;i<n;i++){
            buy[i]=max(buy[i-1],stop[i-1]-prices[i]);
            sell[i]=max(buy[i-1]+prices[i],sell[i-1]);
            stop[i]=max(sell[i-1],stop[i-1]);
        }
        return max(sell[n-1],stop[n-1]);
    }
};
// @lc code=end

