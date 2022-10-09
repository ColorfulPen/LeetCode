package CodeImage.DynamicPrograming;

public class SplitEqualSumSubset {
    /*
    分割相等子集
    解题关键是将求相等解集转化成找到子集内数值和为总数值一半的子集。
    因此转化为01背包问题，即：
    背包容量为总和的一半，物品容量为数值，物品价值也为数值
    若背包放满之后的总价值为总和的一半，那么说明存在该子集，若小于总和的一半，说明不存在这种总和
     */
    public boolean canPartition(int[] nums) {
        int sum = 0;
        for (int i = 0; i < nums.length; i++) {
            sum += nums[i];
        }
        if (sum % 2 != 0) {
            return false;
        }
        int maxWeight = sum / 2;
        int[] dp = new int[maxWeight + 1];
        for (int j = 0; j <= maxWeight; j++) {
            if (nums[0] <= j) {
                dp[j] = nums[0];
            }
        }
        for (int i = 1; i < nums.length; i++) {
            for (int j = maxWeight; j >= 1; j--) {
                if (j >= nums[i]) {
                    dp[j] = Math.max(dp[j], dp[j - nums[i]] + nums[i]);
                }
            }
        }
        if (dp[maxWeight] == maxWeight) {
            return true;
        }
        return false;
    }
}

