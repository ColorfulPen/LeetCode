package CodeImage.DynamicPrograming;

import java.util.Arrays;
import java.util.List;

public class WordBreak {
    public static void main(String[] args) {

    }

    public boolean wordBreak(String s, List<String> wordDict) {
        boolean[] dp = new boolean[s.length() + 1];
        Arrays.fill(dp, false);
        dp[0] = true;
        for (int i = 1; i < s.length() + 1; i++) {
            for (String word : wordDict) {
                if (i >= word.length() && s.startsWith(word, i - word.length()) && dp[i - word.length()]) {
                    dp[i] = true;
                    break;
                }
            }
        }
        return dp[s.length()];
    }
}
