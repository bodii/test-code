// 显示读入的整数值与1、2...的差值

import java.util.Scanner;

class For2Var {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("整数值: ");
		int n = stdIn.nextInt();

		for (int i = 1, j = n -1; i <= n; i++, j--)
			System.out.println(i + " " +j);
	}
}
