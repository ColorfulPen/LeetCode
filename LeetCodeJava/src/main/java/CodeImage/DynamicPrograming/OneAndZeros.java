package CodeImage.DynamicPrograming;

public class OneAndZeros {
    public static void main(String[] args) {

    }

    /*
    其实就是将一维的背包转化成了二维的背包
    dp[i][j]代表i个0，j个1是最大的子集数量，初始化为0
    dp[i][j]=max(dp[i][j],dp[i-zeroNum][j-oneNum]+1)
     */
    public static int getMaxSub(String[] strs, int m, int n) {
        int[][] dp = new int[m + 1][n + 1];

        for (String str : strs) {
//            预处理
            int zeroNums = 0, oneNums = 0;
            for (int i = 0; i < str.length(); i++) {
                if (str.charAt(i) == '0') {
                    zeroNums++;
                } else {
                    oneNums++;
                }
            }

            for (int zeroIndex = m; zeroIndex >= 0; zeroIndex--) {
                for (int oneIndex = n; oneIndex >= 0; oneIndex--) {
                    if (zeroIndex < zeroNums || oneIndex < oneNums) {
                        break;
                    }
                    dp[zeroIndex][oneIndex] = Math.max(dp[zeroIndex][oneIndex], dp[zeroIndex - zeroNums][oneIndex - oneNums] + 1);
                }
            }
        }
        return dp[m][n];
    }
}
