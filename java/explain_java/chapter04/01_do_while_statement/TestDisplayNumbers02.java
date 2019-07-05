// 读入两个整数值，将位于这两个数值之间的所有整数(包括这两个数值)按从小到大的顺序显示出来

import java.util.Scanner;

class TestDisplayNumbers02 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("输入整数A: "); int a = stdIn.nextInt();
		System.out.print("输入整数B: "); int b = stdIn.nextInt();

		do {
			if (a < b) {
				System.out.print(a + " ");
				a ++;
			} else {
				System.out.print(b + " ");
				b ++;
			}
		} while (a != b);

		System.out.print(b + " ");

		System.out.println("");
	}
} 
