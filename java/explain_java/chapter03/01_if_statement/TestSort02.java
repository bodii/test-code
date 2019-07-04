// 读入两个整数值按降序(从大到小的顺序)进行排序

import java.util.Scanner;

class TestSort02 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("整数a: "); int a = stdIn.nextInt();
		System.out.print("整数b: "); int b = stdIn.nextInt();

		int t;
		if (a < b) {
			t = a;
			a = b;
			b = t;
		}

		System.out.println("第一个值: " + a + ", 第二个值: " + b);
	}
}
