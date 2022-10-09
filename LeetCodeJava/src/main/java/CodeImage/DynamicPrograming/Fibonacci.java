package CodeImage.DynamicPrograming;

public class Fibonacci {
    public static void main(String[] args) {
        System.out.println(fib(2));
        System.out.println(fib(3));
        System.out.println(fib(4));
        System.out.println(fib(5));
    }

    public static int fib(int n) {
        if (n == 0) {
            return 0;
        } else if (n == 1) {
            return 1;
        }

        int[] res = new int[n + 1];
        res[0] = 0;
        res[1] = 1;
        for (int index = 2; index < n + 1; index++) {
            res[index] = res[index - 1] + res[index - 2];
        }
        return res[n];
    }

}
