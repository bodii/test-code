// 显示九九乘法表

class MultiTable {
	public static void main(String[] args) {
		for (int i = 1; i <= 9; i++) {
			for (int j = 1; j <= 9; j++) {
				if (i * j < 10)
					System.out.print("  "); // 两个空格
				else
					System.out.print(" "); // 一个空格
				System.out.print(i * j);
			}
			System.out.println();
		}
	}
}
