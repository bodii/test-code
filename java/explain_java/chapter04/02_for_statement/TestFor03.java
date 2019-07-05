// 用for语句来实现计算1到n的和的代码

import java.util.Scanner;

class TestFor03 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("正整数: "); int n = stdIn.nextInt();

		int sum = 0;
		for (int i = 1; i <= n; i++)
			sum += i;
		System.out.println("1到" + n + "之间数的和" + sum);
	}
}
