// 显示直角在左下方的三角形

import java.util.Scanner;

class IsoscelesTriangle {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("显示直角在左下方的三角形。");
		System.out.print("层数: "); int n = stdIn.nextInt();

		for (int i = 1; i <= n; i++) {
			for (int j = 1; j <= i; j++)
				System.out.print('*');
			System.out.println();
		}
	}
}
