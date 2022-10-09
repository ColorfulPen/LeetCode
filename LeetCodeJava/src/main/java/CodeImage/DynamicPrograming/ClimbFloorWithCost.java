package CodeImage.DynamicPrograming;

public class ClimbFloorWithCost {
}
//原始理解
//dp[i]：代表爬到第i层的最小花费
//dp[i]=min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])
//dp[0]=0;dp[1]=cost[1];

//实际理解
//dp[i]：代表爬到第i层的最小花费
//dp[i]=min(dp[i-1],dp[i-2]）+cost[i]
//dp[0]=cost[0];dp[1]=cost[1];