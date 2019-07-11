// 判断两个数组是否相等

import java.util.Scanner;

class ArrayEqual {

	// -- 两个数组a、b中的全部元素是否相等 -- //
	static boolean equals(int[] a, int[] b) {
		if (a.length != b.length)
			return false;

		for(int i = 0; i < a.length; i++) 
			if (a[i] != b[i])
				return false;

		return true;
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("数组a中的元素个数: ");
		int na = stdIn.nextInt();

		int[] a = new int[na];
		for (int i = 0; i < na; i++) {
			System.out.printf("a[%d] : ", i); a[i] = stdIn.nextInt();
		}

		System.out.print("数组b中的元素个数: ");
		int nb = stdIn.nextInt();

		int[] b = new int[nb];
		for (int i = 0; i < nb; i++) {
			System.out.printf("b[%d] : ", i); b[i] = stdIn.nextInt();
		}

		System.out.println(
				"数组a和b" + (
					equals(a, b) ? "相等" : "不相等"
				) 
			);
	}
}
