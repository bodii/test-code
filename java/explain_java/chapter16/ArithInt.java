// 读入两个整数值，并显示加减乘除运算的结果

import java.util.Scanner;

class ArithInt {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("对x和y进行加减乘除运算。");

		System.out.print("x的值: "); int x = stdIn.nextInt();
		System.out.print("y的值: "); int y = stdIn.nextInt();
		
		System.out.println("x + y = " + (x + y));
		System.out.println("x - y = " + (x - y));
		System.out.println("x * y = " + (x * y));
		System.out.println("x / y = " + (x / y));
		System.out.println("x % y = " + (x % y));
	}
}
