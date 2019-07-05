// 显示读入的整数值的所有约数，在显示完约数之后，显示约数的个数

import java.util.Scanner;

class TestFor06 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("正整数: "); int n = stdIn.nextInt();

		int j = 0;
		for (int i = 1; i <= n; i++) {
			if (n % i == 0) {
				System.out.println(i);
				j++;
			}
		}
		System.out.println("约数有" + j + "个。");
	}
}
