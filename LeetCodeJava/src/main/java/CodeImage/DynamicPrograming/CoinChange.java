package CodeImage.DynamicPrograming;

import java.util.Arrays;

public class CoinChange {
    public static void main(String[] args) {
        int[] coins = new int[]{1, 2, 5};
        int[] smallcoins = new int[]{2};
        System.out.println(coinChange(coins, 11));
    }

    public static int coinChange(int[] coins, int amount) {
        int[] dp = new int[amount + 1];
        Arrays.fill(dp, 100000);
        dp[0] = 0;
        Arrays.sort(coins);
        if (amount == 0) {
            return 0;
        }

//        for (Integer coin:coins){
//            for (int i=coin;i<amount+1;i++){
//                if (dp[i]==-1&&dp[i-coin]!=-1){
//                    dp[i]=dp[i-coin]+1;
//                }else {
//                    if (dp[i-coin]!=-1){
//                        dp[i]=Math.min(dp[i],dp[i-coin]+1);
//                    }
//                }
//            }
//        }

        for (Integer coin : coins) {
            for (int i = coin; i < amount + 1; i++) {
                dp[i] = Math.min(dp[i], dp[i - coin] + 1);
            }
        }
        return dp[amount] == 100000 ? -1 : dp[amount];
    }
}

