// 对诗入的整数数组进行加法运算(5个整数 x 10组)

import java.util.Scanner;

class SumGroup2 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("对整数进行加法运算。");
		int total = 0;

	Outer:
		for (int i = 1; i <= 10; i++) {
			System.out.println("第" + i + "组");
			for (int j = 0; j <= 5; j++) {
				System.out.println("整数: ");
				int t = stdIn.nextInt();
				if (t == 99999) 
					break Outer;
				else if (t == 88888)
					continue Outer;
				total += t;
			}
		}

		System.out.println("\n合计值为: " + total + "。");
	}
}
