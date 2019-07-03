// 计算三个整数中的最大值

import java.util.Scanner;

class Max03 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("整数a: "); int a = stdIn.nextInt();
		System.out.print("整数b: "); int b = stdIn.nextInt();
		System.out.print("整数c: "); int c = stdIn.nextInt();

		int max = a;
		if (b > max) max = b;
		if (c > max) max = c;

		System.out.println("最大值是" + max + "。");
	}
}
