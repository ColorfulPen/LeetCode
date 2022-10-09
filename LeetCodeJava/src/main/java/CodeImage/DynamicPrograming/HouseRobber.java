package CodeImage.DynamicPrograming;

public class HouseRobber {
    public static void main(String[] args) {

    }

    /*
    对于循环的数组来说，就是去头和去尾两种情况分别考虑，再比较大小就可以了
    因为不能同时包含头元素和尾元素
     */
    public static int rob(int[] nums) {
        if (nums.length == 1) {
            return nums[0];
        }
        int[] dp = new int[nums.length];
        dp[0] = nums[0];
        dp[1] = Math.max(nums[0], nums[1]);

        for (int i = 2; i < nums.length; i++) {
            dp[i] = Math.max(dp[i - 1], dp[i - 2] + nums[i]);
        }
        return dp[nums.length - 1];
    }
}
