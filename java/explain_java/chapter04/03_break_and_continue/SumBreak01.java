// 对读入的整数进行加法运算(输入0的话就结束)

import java.util.Scanner;

class SumBreak01 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("对整数进行加法运算。");
		System.out.print("要相加的多少个整数: ");
		int n = stdIn.nextInt();

		int sum = 0;
		for (int i = 1; i <= n; i++) {
			System.out.print("整数(以0结束): ");
			int t = stdIn.nextInt();

			if (t == 0) break;  // 跳出for循环
			sum += t;
		}

		System.out.println("合计值为: " + sum + "。");
	}
}
