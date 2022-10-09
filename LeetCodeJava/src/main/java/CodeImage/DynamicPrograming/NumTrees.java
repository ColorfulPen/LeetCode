package CodeImage.DynamicPrograming;

import java.util.Arrays;

public class NumTrees {
    public static void main(String[] args) {

    }

    /*
    dp[i]代表i个节点的二叉搜索树
    dp[i]需要遍历获得
     */
    public static int numTrees(int n) {
        if (n == 1) {
            return 1;
        }
        int[] dp = new int[n + 1];
        Arrays.fill(dp, 0);
        dp[0] = 1;
        dp[1] = 1;
        for (int i = 2; i < n + 1; i++) {
            for (int j = 1; j <= i; j++) {
                dp[i] += dp[j - 1] * dp[i - j];
            }
        }
        return dp[n];
//        dp[2]=2;
//        dp[3] dp[0] dp[2]  dp[1] dp[1] dp[2] dp[0]
    }
}
