package CodeImage.MonotoneStack;

import java.util.Arrays;
import java.util.Stack;

public class pickUpRainWater {
    public static void main(String[] args) {
        int[] array = {0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1};
        int[] test = {3, 2, 1, 2, 3};
//        int[] res = doublePoint(array);
        int[] res = dpSolution(array);
        int result = 0;
        for (int i = 1; i < res.length - 1; i++) {
            System.out.print(res[i] + " ");
            result += res[i];
        }
        System.out.println();
        System.out.println(result);
        System.out.println();
        System.out.println(trap(array));
        System.out.println(singleStack(test));

    }

    public static int[] doublePoint(int[] array) {
        int[] res = new int[array.length];
        Arrays.fill(res, 0);
        int leftHeight, rightHeight;
        for (int i = 1; i < array.length - 1; i++) {
            leftHeight = 0;
            rightHeight = 0;
            for (int leftIndex = 0; leftIndex < i; leftIndex++) {
                leftHeight = Math.max(leftHeight, array[leftIndex]);
            }
            for (int rightIndex = i + 1; rightIndex < array.length; rightIndex++) {
                rightHeight = Math.max(rightHeight, array[rightIndex]);
            }
            res[i] = Math.max(Math.min(leftHeight, rightHeight) - array[i], 0);
        }
        return res;
    }

    public static int[] dpSolution(int[] array) {
        int[] maxLeft = new int[array.length];
        int[] maxRight = new int[array.length];
        int[] res = new int[array.length];
        Arrays.fill(maxLeft, 0);
        Arrays.fill(maxRight, 0);
        for (int i = 1; i < maxLeft.length - 1; i++) {
            maxLeft[i] = Math.max(maxLeft[i - 1], array[i]);
        }
        for (int i = maxRight.length - 2; i > 0; i--) {
            maxRight[i] = Math.max(maxRight[i + 1], array[i]);
        }
        for (int i = 1; i < array.length - 1; i++) {
            res[i] = Math.max(Math.min(maxLeft[i], maxRight[i]) - array[i], 0);
        }
        return res;
    }

    public static int singleStack(int[] array) {
        int res = 0;
        Stack<Integer> stack = new Stack<>();
        stack.push(0);
        for (int i = 1; i < array.length; i++) {
            int stackTop = stack.peek();
            if (array[i] > array[stackTop]) {
                while (!stack.isEmpty() && array[i] > array[stackTop]) {
                    int mid = stack.pop();
                    if (!stack.empty()) {
                        int h = Math.min(array[stack.peek()], array[i]) - array[mid];
                        int w = i - stack.peek() - 1;
                        if (h * w > 0) {
                            res += h * w;
                        }
                        stackTop = stack.peek();
                    }
                }
            } else if (array[i] == array[stackTop]) {
                stack.pop();
            }
            stack.push(i);
        }
        return res;
    }

    public static int trap(int[] height) {
        int size = height.length;

        if (size <= 2) return 0;

        // in the stack, we push the index of array
        // using height[] to access the real height
        Stack<Integer> stack = new Stack<Integer>();
        stack.push(0);

        int sum = 0;
        for (int index = 1; index < size; index++) {
            int stackTop = stack.peek();
            if (height[index] < height[stackTop]) {
                stack.push(index);
            } else if (height[index] == height[stackTop]) {
                // 因为相等的相邻墙，左边一个是不可能存放雨水的，所以pop左边的index, push当前的index
                stack.pop();
                stack.push(index);
            } else {
                //pop up all lower value
                int heightAtIdx = height[index];
                while (!stack.isEmpty() && (heightAtIdx > height[stackTop])) {
                    int mid = stack.pop();

                    if (!stack.isEmpty()) {
                        int left = stack.peek();

                        int h = Math.min(height[left], height[index]) - height[mid];
                        int w = index - left - 1;
                        int hold = h * w;
                        if (hold > 0) sum += hold;
                        stackTop = stack.peek();
                    }
                }
                stack.push(index);
            }
        }

        return sum;
    }


}
