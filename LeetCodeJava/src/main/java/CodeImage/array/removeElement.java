package CodeImage.array;

public class removeElement {
    public static void main(String[] args) {
        int[] array = {0, 1, 2, 3, 3, 0, 4, 2};
        int val = 4;
//        System.out.println(violence(array,val));
//        System.out.println(fastAndSlow(array,val));
        System.out.println(doublePointer(array, val));
    }

    public static int violence(int[] nums, int val) {
        int size = nums.length;
        for (int i = 0; i < size; i++) {
            if (nums[i] == val) {
                if (size - (i + 1) >= 0) System.arraycopy(nums, i + 1, nums, i + 1 - 1, size - (i + 1));
                i--;
                size--;
            }
        }
        return size;
    }

    public static int fastAndSlow(int[] nums, int val) {
        int fastIndex, slowIndex = 0;
        for (fastIndex = 0; fastIndex < nums.length; fastIndex++) {
            if (nums[fastIndex] != val) {
                nums[slowIndex] = nums[fastIndex];
                slowIndex++;
            }
        }
        return slowIndex;
    }

    //    相向双指针法
    public static int doublePointer(int[] nums, int val) {
        int left = 0, right = nums.length - 1;
        while (right >= 0 && right == val) right--;
        while (left <= right) {
            if (nums[left] == val) {
                nums[left] = nums[right];
                right--;
            }
            left++;
            while (right >= 0 && nums[right] == val) right--;
        }
        return left;
    }

}
