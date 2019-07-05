// 直角在左上方、右下方、右上方的三角形程序

import java.util.Scanner;

class TestIsoscelesTriangle01 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("三角形的层数: "); int n = stdIn.nextInt();

		System.out.println("显示直角在左上方的三角形。");
		for (int i = n; i > 0; i--) {
			for (int j = i; j > 0; j--) {
				System.out.print('*');
			}
			System.out.println();
		}

		System.out.println("显示直角在右下方的三角形。");
		for (int i = n; i > 0; i--) {
			for (int s = i; s > 0; s--) {
				System.out.print(' ');
			}
			for (int j = i; j <= n; j++) {
				System.out.print('*');
			}
			System.out.println();
		}

		System.out.println("显示直角在右上方的三角形。");
		for (int i = n; i > 0; i--) {
			for (int s = 0; s < n - i; s++) {
				System.out.print(' ');
			}
			for (int j = i; j > 0; j--) {
				System.out.print('*');
			}
			System.out.println();
		}
	}
}
