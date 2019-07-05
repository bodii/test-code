// 显示读入的整数值的所有约数

import java.util.Scanner;

class Measure {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("整数值: ");
		int n = stdIn.nextInt();

		for (int i = 1; i <= n; i++) 
			if (n % i == 0) // 如果能整除
				System.out.println(i);  // 显示
	}
}
