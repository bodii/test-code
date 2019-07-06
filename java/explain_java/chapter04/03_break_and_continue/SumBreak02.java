// 对读入的整数进行加法运算(在合计值不超过1000的范围内进行加法运算)

import java.util.Scanner;

class SumBreak02 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("对整数进行加法运算。");
		System.out.print("要相加多少个整数: ");
		int n = stdIn.nextInt();

		int sum = 0;
		for (int i = 1; i <= n; i++) {
			System.out.print("整数: ");
			int t = stdIn.nextInt();
			if (sum + t > 1000) {
				System.out.println("合计值超过了1000.");
				System.out.println("最后一个数值被忽略。");
				break;
			}
			sum += t;
		}

		System.out.println("合计值为" + sum + "。");
	}
}
