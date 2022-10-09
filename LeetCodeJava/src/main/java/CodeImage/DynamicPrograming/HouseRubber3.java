package CodeImage.DynamicPrograming;

public class HouseRubber3 {
    public class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode() {
        }

        TreeNode(int val) {
            this.val = val;
        }

        TreeNode(int val, TreeNode left, TreeNode right) {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }

    public static int rob(TreeNode root) {
        int[] dp = caluTotalMoney(root);
        return Math.max(dp[0], dp[1]);

    }

    public static int[] caluTotalMoney(TreeNode node) {
        int[] dp = new int[2];
        if (node == null) {
            return new int[]{0, 0};
        }
        int[] left = caluTotalMoney(node.left);
        int[] right = caluTotalMoney(node.right);
//          不偷
        dp[0] = Math.max(left[0], left[1]) + Math.max(right[0], right[1]);
//        偷
        dp[1] = node.val + left[0] + right[0];
        return dp;
    }
}
