// 对2行3列的矩阵进行加法运算

class Matrix {
	public static void main(String[] args) {
		int[][] a = { { 1, 2, 3 }, { 4, 5, 6 } };
		int[][] b = { { 6, 3, 4 }, { 5, 1, 2 } };
		int[][] c = { { 0, 0, 0 }, { 0, 0, 0 } };

		for (int i = 0; i < 2; i++)
			for (int j = 0; j < 3; j++)
				c[i][j] = a[i][j] + b[i][j];

		System.out.println("矩阵a");
		for (int i = 0; i < 2; i++) {
			for (int j = 0; j < 3; j++) 
				System.out.printf("%3d", a[i][j]);
			System.out.println();
		}

		System.out.println("矩阵b");
		for (int i = 0; i < 2; i++) {
			for (int j = 0; j < 3; j++) 
				System.out.printf("%3d", b[i][j]);
			System.out.println();
		}

		System.out.println("矩阵c");
		for (int i = 0; i < 2; i++) {
			for (int j = 0; j < 3; j++)
				System.out.printf("%3d", c[i][j]);
			System.out.println();
		}

	}
}
