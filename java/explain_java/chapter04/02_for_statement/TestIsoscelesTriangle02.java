// 显示n层的金字塔
// 第i行显示(i - 1) * 2 + 1个'*',最后一行即第n行显示（n - 1) * 2 + 1个'*'

import java.util.Scanner;

class TestIsoscelesTriangle02 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("几层的三角: "); int n = stdIn.nextInt();

		for (int i = 1; i <= n; i++) {
			for (int s = n - i; s > 0; s--)
				System.out.print(' ');
			for (int j = (i - 1) * 2 + 1; j > 0; j--) 
				System.out.print('*');

			System.out.println();
			
		}
	}
}
