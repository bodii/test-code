// 显示读入的两个整数值中较小的值

import java.util.Scanner;

class Min2 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("整数a: "); int a = stdIn.nextInt();
		System.out.print("整数b: "); int b = stdIn.nextInt();

		int min;  // 较小的值
		if (a < b) 
			min = a;
		else 
			min = b;

		System.out.println("较小的值是" + min + "。");
	}
}
