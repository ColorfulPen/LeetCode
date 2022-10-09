package CodeImage.DynamicPrograming;

public class IntegerSpilt {
    public static void main(String[] args) {
        for (int i = 2; i <= 10; i++) {
            System.out.println(integerSpilt(i));
        }
    }
/*
dp[i]表示正整数i所拆分的乘积最大化
dp[i]=dp[j]*dp[i-j];  j>=2 and j<=i/2

4 2,2
5 2,3
6 2,4   3,3
7 2,5   3,4
8 2,6 3,5 4,4
9 2,7 3,6 4,5
10 2,8 3,7 4,6 5,5
dp[2]=1,dp[3]=2,dp[4]=max(2*2,dp[2]*dp[2])=4,dp[5]=max(2*3,dp[2]*dp[3])=6,
dp[6]=max(3*3,dp[3]*dp[3],dp[2]*dp[4])=9
dp[7]=max(3*4,)=12
dp[8]=max(4*4,)=16
dp[9]=max(4*5,dp[4]*dp[5])=24
dp[10]=max(5*5,)
这里存在一个问题，对于dp[2]*dp[x]来说，dp[2]可以作为2使用，而不是1，因此会导致出错
同理
修改：把dp[2]和dp[3]设为最大值即可
 */

    /*
    思路二：
    dp[i]=max(j*dp[i-j],dp[j]*dp[i-j])
     */
    public static int integerSpilt(int n) {
        if (n == 2) {
            return 1;
        } else if (n == 3) {
            return 2;
        }
        int[] dp = new int[n + 1];
        dp[2] = 2;
        dp[3] = 3;
        for (int i = 4; i < n + 1; i++) {
            dp[i] = (i / 2) * (i - i / 2);
            for (int j = 2; j <= i / 2; j++) {
                dp[i] = Math.max(dp[i], dp[j] * dp[i - j]);
            }
        }
        return dp[n];
    }
}
