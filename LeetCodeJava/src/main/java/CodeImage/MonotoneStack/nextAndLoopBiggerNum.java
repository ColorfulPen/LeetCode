package CodeImage.MonotoneStack;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Stack;

public class nextAndLoopBiggerNum {
    public static void main(String[] args) {
        int[] nums1 = {4, 1, 2};
        int[] nums2 = {1, 3, 4, 2};
        int[] nums3 = {1, 3, 5, 2, 4};
        int[] nums4 = {5, 4, 3, 2, 1};

        for (int i : nextBig(nums3, nums4)) {
            System.out.println(i + " ");
        }

    }

    public static int[] nextBig(int[] nums1, int[] nums2) {
        int[] res = new int[nums1.length];
        Arrays.fill(res, -1);
        Stack<Integer> stack = new Stack<>();
        HashMap<Integer, Integer> hashMap = new HashMap<>();
        for (int i = 0; i < nums1.length; i++) {
            hashMap.put(nums1[i], i);
        }
        stack.push(0);
        for (int i = 1; i < nums2.length; i++) {
            while (!stack.isEmpty() && nums2[stack.peek()] < nums2[i]) {
                int prev = stack.pop();
                if (hashMap.containsKey(nums2[prev])) {
                    res[hashMap.get(nums2[prev])] = nums2[i];
                }
            }
            stack.push(i);
        }
        return res;
    }
}
