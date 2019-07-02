// 读入两个整数值，并显示加减乘除运算的结果

import java.util.Scanner;

class ArithInt {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("对x和y进行加减乘除运算。");
		
		System.out.print("x的值: ");     // 提示输入x的值
		int x = stdIn.nextInt();         // 读入x的整数值

		System.out.print("y的值: ");     // 提示输入y的值
		int y = stdIn.nextInt();         // 读入y的整数值

		System.out.println("x + y = " + (x + y));   // 显示 x + y 的值
		System.out.println("x - y = " + (x - y));   // 显示 x - y 的值
		System.out.println("x * y = " + (x * y));   // 显示 x * y 的值
		System.out.println("y / y = " + (x / y));   // 显示 x / y 的值(商)
		System.out.println("x % y = " + (x % y));   // 显示 x % y 的值(余数)
	}
}
