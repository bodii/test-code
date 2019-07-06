// 对读入的整数进行加法运算、平均值(不对负值运算)

import java.util.Scanner;

class TestSumContinue01 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("对整数进行加法运算、平均值运算。");
		System.out.print("要相加多少个整数: "); int n = stdIn.nextInt();

		int sum = 0;
		for (int i = 1; i <= n; i++) {
			System.out.print("整数: "); int t = stdIn.nextInt();
			if (t < 0) {
				System.out.println("不对负值进行加法运算。");
				continue;
			}
			sum += t;
		}

		System.out.println("合计值为: " + sum + "。");
		System.out.println("平均值为: " + sum / 2 + "。");
	}
}
