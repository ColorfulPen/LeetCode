package CodeImage.DynamicPrograming;

public class CombinationSumFour {
    public static void main(String[] args) {
        int[] nums = new int[]{1, 2, 3};
        int target = 4;
    }

    /*
    先遍历容量  再遍历物品
    dp[0] dp[1] dp[2] dp[3] dp[4]
      1     1     2     4     7
     */
    public static int combinationSum4(int[] nums, int target) {
        int[] dp = new int[target + 1];
        dp[0] = 1;
        for (int weight = 1; weight < target + 1; weight++) {
            for (Integer n : nums) {
                if (weight >= n) {
                    dp[weight] += dp[weight - n];
                }
            }
        }
        return dp[target];
    }
}
