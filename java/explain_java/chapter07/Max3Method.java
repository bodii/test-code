// 计算3个整数值中的最大值

import java.util.Scanner;

class Max3Method {
	// --- 返回a、b、c中的最大值 --- //
	static int max(int a, int b, int c) {
		int max = a;
		if (b > max) max = b;
		if (c > max) max = c;

		return max;
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("整数a: "); int a = stdIn.nextInt();
		System.out.print("整数b: "); int b = stdIn.nextInt();
		System.out.print("整数c: "); int c = stdIn.nextInt();

		System.out.println("最大值是: " + max(a, b, c) + "。");
	}
}
