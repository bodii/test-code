// 计算两个数值中的最大值、三个数值中的最大值的方法(重载)

import java.util.Scanner;

class Max {

	// -- 返回a、b中的最大值 -- //
	static int max(int a, int b) {
		return a > b ? a : b;
	}

	// -- 返回a、b、c中的最大值 --//
	static int max(int a, int b, int c) {
		int max = a;
		if (b > max) max = b;
		if (c > max) max = c;
		return max;
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("x的值: "); int x = stdIn.nextInt();
		System.out.print("y的值: "); int y = stdIn.nextInt();
		System.out.print("z的值: "); int z = stdIn.nextInt();

		// 两个值中的最大值
		System.out.println("x, y中的最大值是" + max(x, y) + "。");

		// 三个值中的最大值
		System.out.println("x, y, z中的最大值是" + max(x, y, z) + "。");
	}
}
