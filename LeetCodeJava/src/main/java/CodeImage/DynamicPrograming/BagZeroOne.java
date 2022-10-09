package CodeImage.DynamicPrograming;

public class BagZeroOne {
    public static void main(String[] args) {
        int[] weight = new int[]{1, 3, 4};
        int[] value = new int[]{15, 20, 30};
        int[] nums = new int[]{1, 5, 9, 7};
        System.out.println(bag01(weight, value, 4));
        System.out.println(bag01OneDimension(weight, value, 4));
        System.out.println(bag01(nums, nums, 11));
        System.out.println(bag01OneDimension(nums, nums, 11));
        bag01OneDimension(nums, nums, 11);
    }

    /*
    dp[i][j]代表容量为j的情况下，从0-i件物品中选，价值最大的情况
    dp[i][j]=max(dp[i-1][j],dp[i-1][j-w[i]]+v[i])
    解释：dp[i-1][j]含义就是不放i物品进去，为啥不放？因为放不下
    dp[i-1][j-w[i]]+v[i]含义是放i进去后的最大价值
     */
    public static int bag01(int[] weight, int[] value, int maxWeight) {
        int itemSize = weight.length;
        int[][] dp = new int[itemSize][maxWeight + 1];
//        初始化行（价值），列无需初始化，因为默认为0
        for (int j = 0; j < maxWeight + 1; j++) {
            if (weight[0] <= j) {
                dp[0][j] = value[0];
            }
        }
//        递推，从物品或者重量递推都可以
        for (int i = 1; i < itemSize; i++) {
            for (int j = 1; j < maxWeight + 1; j++) {
                if (j < weight[i]) {
                    dp[i][j] = dp[i - 1][j];
                } else {
                    dp[i][j] = Math.max(dp[i - 1][j], dp[i - 1][j - weight[i]] + value[i]);
                }
            }
        }

        return dp[itemSize - 1][maxWeight];
    }

    //    ！！！！！！！！！！！！！！！！！！！！
//    ！！！！！！！！！！！！！！！！！！！！
//    对于一维数组的递推，要切记从数组大值递推到小值，因为若从小值推导到大值，可能会导致推导到大值时使用最新的小index的值
//    这点需要注意
    public static int bag01OneDimension(int[] weight, int[] value, int maxWeight) {
        int[] dp = new int[maxWeight + 1];
//        初始化
        for (int j = 0; j <= maxWeight; j++) {
            if (weight[0] <= j) {
                dp[j] = value[0];
            }
        }
//        递推
        for (int i = 1; i < weight.length; i++) {
//            for (int j = 1; j < maxWeight + 1; j++) {
//                if (j >= weight[i]) {
//                    dp[j] = Math.max(dp[j], dp[j - weight[i]] + value[i]);
//                }
//            }
            for (int j = maxWeight; j >= 1; j--) {
                if (j >= weight[i]) {
                    dp[j] = Math.max(dp[j], dp[j - weight[i]] + value[i]);
                }
            }
        }
        return dp[maxWeight];
    }
}
