import java.util.Random;
import java.util.Scanner;

class Abc {
	public static void main(String[] args) {
		Random rand = new Random();
		Scanner stdIn = new Scanner(System.in);

		int a;        // a是int型变量
		a = 2;        // 赋值（将值赋给创建好的变量)
		int b = -1;   // 初始化(创建变量时赋入值)
		double x = 1.5 * 2; 
		final double PI = 3.14;
		x = rand.nextDouble();

		System.out.println(
			"半径为" + x + "的圆的面积是" + (PI * x * x) + "。"
		);

		System.out.print("整数a的值: ");
		a = stdIn.nextInt();
		System.out.println("a / 2 = " + a / 2);
		System.out.println("a % 2 = " + a % 2);

		// 字符串类型
		String s = "ABC";
		System.out.println("字符串s是" + s + "。");

	}
}
