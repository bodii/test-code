// 以3个数组x、y、z的行数和列数都相同为前提。编写一个方法，如果3个数组的元素个数相等，则
// 执行加法运算，并返回true,如果不相等，则不执行加法运算，返回false

class TestAddMatrix01 {
	static boolean addMatrix(int[][] x, int[][] y, int[][] z) {
		if (x.length != y.length || x.length != z.length)
			return false;

		for (int i = 0; i < x.length; i++) 
			if (x[i].length != y[i].length || x[i].length != z[i].length)
				return false;
		
		return true;
	}

	static int[][] sum(int[][] x, int[][] y, int[][] z) {
		int[][] s = new int[x.length][x[0].length];
		for (int i = 0; i < x.length; i++)
			for (int j = 0; j < x[i].length; j++)
				s[i][j] = x[i][j] + y[i][j] + z[i][j];

		return s;
	}

	static void printMatrix(int[][] m) {
		for (int i = 0; i < m.length; i++) {
			for (int j = 0; j < m[i].length; j++)
				System.out.printf("%-3d", m[i][j]);
			System.out.println();
		}
		System.out.println();
	}

	public static void main(String[] args) {
		int[][] x = { { 1, 2, 3 }, { 4, 5, 6 } };
		int[][] y = { { 6, 3, 4 }, { 5, 1, 2 } };
		int[][] z = { { 3, 8, 5 }, { 9, 3, 7 } };

		System.out.println("矩阵x");
		printMatrix(x);

		System.out.println("矩阵y");
		printMatrix(y);

		System.out.println("矩阵z");
		printMatrix(z);

		System.out.println("三个矩阵个元素相加:");
		if (addMatrix(x, y ,z)) {
			int[][] s = sum(x, y, z);
			printMatrix(s);
		}



	}
}
