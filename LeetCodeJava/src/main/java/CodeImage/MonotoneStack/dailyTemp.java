package CodeImage.MonotoneStack;

import java.util.Stack;

public class dailyTemp {
    public static void main(String[] args) {
        int[] temperature = {73, 74, 75, 71, 69, 72, 76, 73};
        int[] res = getHigh(temperature);
        for (int r : res) {
            System.out.print(r + " ");
        }
    }

    public static int[] getHigh(int[] temperature) {
        Stack<Integer> stack = new Stack<>();
        int[] res = new int[temperature.length];
        stack.push(0);
        for (int i = 1; i < temperature.length; i++) {
            while (!stack.isEmpty() && temperature[stack.peek()] < temperature[i]) {
                int prev = stack.pop();
                res[prev] = i - prev;
            }
            stack.push(i);

        }
        return res;
    }
}
