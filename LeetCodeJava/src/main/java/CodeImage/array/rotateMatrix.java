package CodeImage.array;

public class rotateMatrix {
    public static void main(String[] args) {
        int n = 7;
        printMatrix(generateMatrix(n));

    }

    public static int[][] generateMatrix(int n) {
        int[][] matrix = new int[n][n];
        int startX = 0, startY = 0;
        int loop = n / 2;
        int mid = n / 2;
        int count = 1;
        int offset = 1;
        int i, j;
        while (loop > 0) {
            i = startX;
            j = startY;
            for (j = startY; j < n - offset; j++) {
                matrix[startX][j] = count++;
            }
            for (i = startX; i < n - offset; i++) {
                matrix[i][j] = count++;
            }
            // 模拟填充下行从右到左(左闭右开)
            for (; j > startY; j--) {
                matrix[i][j] = count++;
            }
            // 模拟填充左列从下到上(左闭右开)
            for (; i > startX; i--) {
                matrix[i][j] = count++;
            }
            startX++;
            startY++;
            offset++;
            loop--;
        }
        if (n % 2 != 0) {
            matrix[mid][mid] = count;
        }
        return matrix;
    }

    public static void printMatrix(int[][] matrix) {
        for (int[] ints : matrix) {
            for (int j = 0; j < matrix[0].length; j++) {
                System.out.print(ints[j] + " ");
            }
            System.out.println();
        }
    }
}
