package CodeImage.DynamicPrograming;

public class ClimbFloor {
    public static void main(String[] args) {

    }

//    dp[i]：表示爬到第i个阶梯有几种方法
//    dp[i]=dp[i-1]+dp[i-2]
//    dp[0]=1,dp[1]=1;

    /*
    使用完全背包求解
    dp[i]代表爬到第i个阶梯有多少种方法
    dp[0]=1
     */
    public static int climb(int n) {
        int[] dp = new int[n + 1];
        int[] feets = new int[]{1, 2};
        dp[0] = 1;
        for (int i = 1; i < n + 1; i++) {
            for (Integer feet : feets) {
                if (i >= feet) {
                    dp[i] += dp[i - feet];
                }
            }
        }
        return dp[n];
    }
}
