// 读入的三个整数值按升序(从小到大的顺序)进行排序

import java.util.Scanner;

class TestSort03 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("整数a的值:"); int a = stdIn.nextInt();
		System.out.print("整数b的值:"); int b = stdIn.nextInt();
		System.out.print("整数c的值:"); int c = stdIn.nextInt();

		int t;
		if (a > c && b > a) {
			t = b;
			a = c;
			b = a;
			c = t;
		} else if (a > c && c > b) {
			t = a;
			a = b;
			b = c;
			c = t;
		} else if (a > b && b > c) {
			t = a;
			a = c;
			c = t;
		} else if (a > b && c > a) {
			t = a;
			a = b;
			b = t;
		} else if (a < b && b > c) {
			t = b;
			b = c;
			c = t;
		} else if (c > a && b < a) {
			t = b;
			a = t;
			b = a;
		}

		System.out.println("第一个数: " + a + ", 第二个数: " + b + ", 第三个数: " + c);
	}
}
