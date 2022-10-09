package CodeImage.DynamicPrograming;

public class TargetSum {
    public static void main(String[] args) {
        int[] nums = new int[]{1, 1, 1, 1, 1};
        int ret = getSum(nums, 3);
        System.out.println(ret);
    }
/*
错误分析：
nums[i]代表数据数组
dp[i][j]代表前i个数组合成j的数量
应该要加一个偏移量
problem：
dp[i][j]=dp[i-1][j-nums[i]]+dp[i-1][j+nums[i]]

初始化：
dp[0][j]= nums[i] or
 */

    /*
    正确分析：
    要求：求加减结果为target的个数
    即：left-right=target,又left+right=sum
    因此：left-(sum-left)=target  =>  left=(sum+target)/2
    注意：如果除不尽，说明无法找到left的整数解
    dp[i]代表了结果为i的总个数
    dp[i]就为dp[i-nums[j]]求和所得的值
    dp初始化：
    dp[0]=1
     */
    public static int getSum(int[] nums, int target) {
        int sum = 0;
        for (Integer n : nums) {
            sum += n;
        }
        if ((sum + target) % 2 != 0) {
            return 0;
        }
        int left = Math.abs((sum + target) / 2);
        int[] dp = new int[left + 1];
        dp[0] = 1;
//        for (int i=1;i<left+1;i++){
//            for (Integer n:nums){
//                if (n<=i){
//                    dp[i]+=dp[i-n];
//                }
//            }
//        }
        for (Integer n : nums) {
            for (int i = left; i >= n; i--) {
                dp[i] += dp[i - n];
            }
        }
        return dp[left];
    }
}
