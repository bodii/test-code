// 显示int型一维数组和int二维数组(每行中的列数可能会不一样）中全部元素的值

class TestPrintArray {
	static void printArray(int[] a) {
		for (int i = 0; i < a.length; i++) 
			System.out.printf("%-4d", a[i]);

		System.out.println();
	}

	static void printArray(int[][] a) {
		for (int i = 0; i < a.length; i++) {
			for (int j = 0; j < a[i].length; j++)
				System.out.printf("%-7d", a[i][j]);
			System.out.println();
		}
		System.out.println();
	}

	public static void main(String[] args) {
		int[] a = { 12, 536, -8, 7 };
		int[][] a1 = { 
			{32, -1, 32, 45, 67},
			{535, 99999, 2},
			{2, 5, -123, 9},
		};

		System.out.println("int型一维数组:");
		printArray(a);
		System.out.println("\nint型二维数组:");
		printArray(a1);

	}
}
