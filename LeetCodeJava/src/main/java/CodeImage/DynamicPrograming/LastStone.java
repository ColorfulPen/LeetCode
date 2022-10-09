package CodeImage.DynamicPrograming;

//1049
public class LastStone {
    public static void main(String[] args) {
        int[] stones = new int[]{2, 7, 4, 1, 8, 1};
        System.out.println(minWeight(stones));
    }

    public static int minWeight(int[] stones) {
        int sum = 0;
        for (int i = 0; i < stones.length; i++) {
            sum += stones[i];
        }
        int target = sum / 2;
        int[] dp = new int[target + 1];

        for (int j = 0; j < target + 1; j++) {
            if (j >= stones[0]) {
                dp[j] = stones[0];
            }
        }

        for (int i = 1; i < stones.length; i++) {
            for (int j = target; j >= 1; j--) {
                if (stones[i] <= j) {
                    dp[j] = Math.max(dp[j], dp[j - stones[i]] + stones[i]);
                }
            }
        }

        int rest = sum - dp[target];

        return Math.abs(rest - dp[target]);
    }
}
