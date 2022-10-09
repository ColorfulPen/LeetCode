package CodeImage.DynamicPrograming;

import java.util.Arrays;

public class PerfectSquares {
    public static void main(String[] args) {
        System.out.println(numSquares(12));
    }

    public static int numSquares(int n) {
        int[] dp = new int[n + 1];
        Arrays.fill(dp, 10000);
        dp[0] = 0;

        int maxNum = (int) Math.sqrt(n);
        for (int i = 1; i <= maxNum; i++) {
            int temp = i * i;
            for (int j = temp; j < n + 1; j++) {
                dp[j] = Math.min(dp[j], dp[j - temp] + 1);
            }
        }
        return dp[n];
    }
}
