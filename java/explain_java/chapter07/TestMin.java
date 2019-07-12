// 编写下述的重载方法群，分别计算两个int型整数a、b中的最小值，三个int型整数a、b、
// c中的最小值，数组a中元素的最小值

import java.util.Scanner;

class TestMin {
	static int min(int a, int b) {
		return a < b ? a : b;
	}

	static int min(int a, int b, int c) {
		int min = a;
		if (min > b)
			min = b;
		if (min > c)
			min = c;

		return min;
	}

	static int min(int[] a) {
		int min = a[0];
		for (int i = 1; i < a.length; i++)
			if (a[i] < min)
				min = a[i];

		return min;
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("a的值: "); int a = stdIn.nextInt();
		System.out.print("b的值: "); int b = stdIn.nextInt();
		System.out.printf("a和b的最小值: %d\n\n", min(a, b));
		
		System.out.print("c的值: "); int c = stdIn.nextInt();
		System.out.printf("a、b、c的最小值: %d\n\n", min(a, b, c));

		System.out.print("数组a多少个元素: "); int n = stdIn.nextInt();
		int[] a1 = new int[n];
		for (int i = 0; i < n; i++) {
			System.out.printf("数组a[%d]的值: ", i); a1[i] = stdIn.nextInt();
		}

		System.out.printf("数组a中元素的最小值: %d\n\n", min(a1));
	}
}
