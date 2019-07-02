/**
 * test
 * 读入两个实数值，求它们的和与平均值并显示结果
 */

import java.util.Scanner;

class ArithDoubleTest01 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("读入两个实数值，求它们的和与平均值并显示结果>>");

		System.out.print("x的值: ");
		double x = stdIn.nextDouble();

		System.out.print("y的值: ");
		double y = stdIn.nextDouble();

		System.out.println("x + y = " + (x + y) + "。");
		System.out.println("(x + y) / 2 = " + (x + y) / 2 + "。");
	}
}
