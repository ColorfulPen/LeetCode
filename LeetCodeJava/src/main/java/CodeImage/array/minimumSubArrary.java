package CodeImage.array;

public class minimumSubArrary {
    public static void main(String[] args) {
        int[] array = {2, 3, 1, 2, 4, 3};
        System.out.println(getMinSubArrary(array, 8));
    }

    public static int getMinSubArrary(int[] arrary, int target) {
        int result = Integer.MAX_VALUE;
        int sum = 0;
        int i = 0;
        for (int j = 0; j < arrary.length; j++) {
            sum += arrary[j];
            while (sum >= target) {
                result = Math.min(result, (j - i + 1));
                sum -= arrary[i++];
            }
        }
        return result == Integer.MAX_VALUE ? 0 : result;
    }
}
