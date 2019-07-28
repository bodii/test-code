// 计算并显示通过命令行参数传入的半径的圆的周长和面积

import java.math.*;
class Circle  {
	public static void main(String[] args) {
		Double r = Double.parseDouble(args[0]);

		System.out.println("圆的周长: " + (r * 2 * Math.PI));
		System.out.println("圆的面积: " + (r * r * Math.PI));
	}
}
