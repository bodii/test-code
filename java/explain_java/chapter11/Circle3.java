// 计算圆的面积（静态导入System.in和System.out)

import java.util.Scanner;
import static java.lang.Math.PI;
import static java.lang.System.in;
import static java.lang.System.out;

class Circle3 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(in);
		out.println("计算圆的面积。");
		out.print("半径: ");
		double r = stdIn.nextDouble();
		out.println("面积为" + (PI * r * r) + "。");
	}
}
