// 读入两个整数值，将位于这两个数值之间的所有整数(包括这两个数值)按从小到大的顺序显示出来

import java.util.Scanner;

class TestDisplayNumbers {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("输入整数A: "); int a = stdIn.nextInt();
		System.out.print("输入整数B: "); int b = stdIn.nextInt();

		int t;
		if (a > b) {
			t = a;
			a = b;
			b = t;
		}

		do {
			System.out.print(a + " ");
			a ++;
		} while (a <= b);

		System.out.println("");
	}
}
