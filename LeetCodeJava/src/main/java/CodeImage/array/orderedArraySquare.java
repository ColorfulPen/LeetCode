package CodeImage.array;

public class orderedArraySquare {
    public static void main(String[] args) {
        int[] array = {-4, -1, 0, 3, 10};
        int[] res = getSquare(array);
        for (int i = 0; i < 5; i++) {
            System.out.println(res[i]);
        }
    }

    public static int[] getSquare(int[] array) {
        int left = 0, right = array.length - 1, count = right;
        int[] res = new int[right + 1];
        while (left <= right) {
            if (array[left] * array[left] <= array[right] * array[right]) {
                res[count--] = array[right] * array[right];
                right--;
            } else {
                res[count--] = array[left] * array[left];
                left++;
            }
        }
        return res;
    }
}
