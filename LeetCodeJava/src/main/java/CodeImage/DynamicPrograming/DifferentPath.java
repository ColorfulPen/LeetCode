package CodeImage.DynamicPrograming;

public class DifferentPath {
}
/*
Path1
dp[i][j]代表走到该格子的种类数
dp[i][j]=dp[i-1][j]+dp[i][j-1]
dp[0][0]=0;dp[0][1]=1;dp[1][0]=1
if i or j <0，置为0
 */

/*
Path2
添加有障碍版本
其实就是将障碍格子的值置为0
 */