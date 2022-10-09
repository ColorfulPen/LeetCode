package CodeImage.array;

public class binarySearch {
    public static void main(String[] args) {
        int[] array = {1, 2, 3, 4, 7, 9, 10};
        System.out.println(search(array, 7));
        System.out.println(search1(array, 10));
    }

    public static int search(int[] nums, int target) {
//        左闭右闭
        int left = 0, right = nums.length - 1;
        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] > target) {
                right = mid - 1;
            } else if (nums[mid] < target) {
                left = mid + 1;
            } else {
                return mid;
            }
        }
        return -1;
    }

    public static int search1(int[] nums, int target) {
//        左闭右开
        int left = 0, right = nums.length;
        while (left < right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] > target) {
                right = mid;
            } else if (nums[mid] < target) {
                left = mid;
            } else {
                return mid;
            }
        }
        return -1;
    }

}

