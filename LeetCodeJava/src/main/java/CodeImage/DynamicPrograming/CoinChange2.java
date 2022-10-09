package CodeImage.DynamicPrograming;

public class CoinChange2 {
    public static void main(String[] args) {
        int[] coins = new int[]{1, 2, 5};
        int amount = 5;
        System.out.println(change(amount, coins));
    }

    /*
    dp[i]代表总价为i的组合方式
    dp[i]=sum(dp[i-coin])
    dp[0]=1
     */
    public static int change(int amount, int[] coins) {
        int[] dp = new int[amount + 1];
        dp[0] = 1;
        for (Integer coin : coins) {
            for (int amountIndex = coin; amountIndex < amount + 1; amountIndex++) {
                dp[amountIndex] += dp[amountIndex - coin];
            }
        }
        return dp[amount];
    }
}
