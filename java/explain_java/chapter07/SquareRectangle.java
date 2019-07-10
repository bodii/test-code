// 显示长方形和正方形

import java.util.Scanner;

class SquareRectangle {
	// -- 连续显示n个字符c -- //
	static void putChars(char c, int n) {
		while (n-- > 0)
			System.out.print(c);
	}

	// -- 通过排列字符'+'来显示边长为n的正方形 -- //
	static void putSquare(int n) {
		for (int i = 1; i <= n; i++) {	// 执行n次下述处理
			putChars('+', n);			//  * 显示n个字符'+'
			System.out.println();		// * 换行
		}
	}

	// -- 通过排列字符'*'来显示宽为h长为w的长方形 -- //
	static void putRectangle(int h, int w) {
		for (int i = 1; i <= h; i++) {		// 执行h次下述处理
			putChars('*', w);				// *显示w个字符'*'
			System.out.println();			// * 换行
		}
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("显示正方形。");
		System.out.print("边长: "); int n = stdIn.nextInt();
		putSquare(n);			// 显示正方形

		System.out.println("显示长方形。");
		System.out.print("宽: "); int h = stdIn.nextInt();
		System.out.print("长: "); int w = stdIn.nextInt();
		putRectangle(h, w);		// 显示长方形

	}
}
