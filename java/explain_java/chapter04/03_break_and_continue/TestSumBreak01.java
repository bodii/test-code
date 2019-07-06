// 读入的整数时行加法运算、平均值(合计值不超过1000的范围内)

import java.util.Scanner;

class TestSumBreak01 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("对整数进行加法运算、平均值。");
		System.out.print("要相加多少个整数: "); int n = stdIn.nextInt();

		int sum = 0;
		for (int i = 0; i < n; i++) {
			System.out.print("正整数; "); int t = stdIn.nextInt();
			if (sum + t > 1000) {
				System.out.println("合计值超过了1000。");
				System.out.println("最后一个数值被忽略。");
				break;
			}

			sum += t;
		}

		System.out.println("合计值为" + sum + "。");
		System.out.println("平均值为" + sum / 2 + "。");
	}
}
