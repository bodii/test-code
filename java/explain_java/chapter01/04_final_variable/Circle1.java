// 计算圆的周长和面积（其1： 使用浮点型常量来表示圆周率)

import java.util.Scanner;

class Circle1 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("半径: ");
		double r = stdIn.nextDouble();

		System.out.println("周长是" + 2 * 3.14 * r + "。");
		System.out.println("面积是" + 3.14 * r * r + "。");
	}
}
