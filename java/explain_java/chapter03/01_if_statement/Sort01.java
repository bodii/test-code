// 将两个变量按升序（从小到大的顺序) 时行排序

import java.util.Scanner;

class Sort01 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);
		
		System.out.print("变量a： "); int a = stdIn.nextInt();
		System.out.print("变量b： "); int b = stdIn.nextInt();

		if (a > b) {
			int t = a;
			a = b;
			b = t;
		}

		System.out.println("排序成a <=b。");
		System.out.println("变量a是" + a + "。");
		System.out.println("变量b是" + b + "。");

	}
}
