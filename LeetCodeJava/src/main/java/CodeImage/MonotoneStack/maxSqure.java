package CodeImage.MonotoneStack;

import java.util.Stack;

public class maxSqure {
    public static void main(String[] args) {
        int[] height = {2, 1, 5, 6, 2, 3};
        int[] height2 = {2, 4};
        int[] height3 = {};
        System.out.println(getMaxSquareDp(height2));
        System.out.println(getMaxSquareSingleStack(height2));
    }

    //    *********************
/*
为了防止数组为单调情况，给数组头尾添加两个0
 */
    public static int getMaxSquareSingleStack(int[] height) {
        int sum = 0;
        int length = height.length;
        if (height.length == 0) {
            return sum;
        }

        int[] newHeight = new int[length + 2];
        newHeight[0] = 0;
        newHeight[newHeight.length - 1] = 0;
        System.arraycopy(height, 0, newHeight, 1, height.length);

        height = newHeight;
        Stack<Integer> stack = new Stack<>();
        stack.push(0);
        for (int i = 1; i < length + 2; i++) {
            int midHeight = height[stack.peek()];
            if (height[i] > midHeight) {
                stack.push(i);
            } else if (height[i] == midHeight) {
                stack.pop();
                stack.push(i);
            } else {
                while (height[i] < height[stack.peek()]) {
                    int mid = stack.peek();
                    stack.pop();
                    int leftIndex = stack.peek();
                    int rightIndex = i;
                    int w = rightIndex - leftIndex - 1;
                    int h = height[mid];
                    sum = Math.max(sum, w * h);
                }
                stack.push(i);
            }
        }
        return sum;
    }


    public static int getMaxSquareDp(int[] height) {
        if (height.length == 0) {
            return 0;
        }

        int length = height.length;
        int[] leftIndex = new int[length];
        int[] rightIndex = new int[length];
        int sum = 0;

//        Find the left first small index
        leftIndex[0] = -1;
        for (int i = 1; i < length; i++) {
            int t = i - 1;
            while (t >= 0 && height[t] >= height[i]) {
                t = leftIndex[t];
            }
            leftIndex[i] = t;
        }

//        Find the right first small index
        rightIndex[length - 1] = length;
        for (int i = length - 2; i >= 0; i--) {
            int t = i + 1;
            while (t < length && height[t] >= height[i]) {
                t = rightIndex[t];
            }
            rightIndex[i] = t;
        }

//        Sum
        for (int i = 0; i < length; i++) {
            int indexHegiht = height[i];
            int indexWidth = rightIndex[i] - leftIndex[i] - 1;
            sum = Math.max(sum, indexHegiht * indexWidth);
        }
        return sum;
    }

}
