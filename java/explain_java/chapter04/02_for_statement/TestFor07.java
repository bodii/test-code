// 显示1到n的整数值的平方

import java.util.Scanner;

class TestFor07 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("正整数: "); int n = stdIn.nextInt();

		for (int i = 1; i <= n; i++) 
			System.out.println(i + "的平方是" + (i * i));
	}
}
