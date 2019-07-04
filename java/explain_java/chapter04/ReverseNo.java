// 逆序显示读入的正整数值

import java.util.Scanner;

class ReverseNo {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("逆序显示正整数值。");
		int x;
		do {
			System.out.print("正整数值: ");
			x = stdIn.nextInt();
		} while (x <= 0);

		System.out.print("倒过来读是: ");
		while (x > 0) {
			System.out.print( x % 10);
			x /= 10;    // x 除以10
		}

		System.out.println("。");
	}
}
