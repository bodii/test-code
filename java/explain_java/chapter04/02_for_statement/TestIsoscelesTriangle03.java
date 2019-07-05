// 显示n层的数字金字塔
// 第i行显示i % 10

import java.util.Scanner;

class TestIsoscelesTriangle03 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("数字金字塔的层数: "); int n = stdIn.nextInt();

		for (int i = 1; i <= n; i++) {
			for (int s = n - i; s > 0; s--) 
				System.out.print(' ');
			for (int j = 1; j <= (i - 1) * 2 + 1; j++) 
				System.out.print(i);
			System.out.println();
		}
	}
}
