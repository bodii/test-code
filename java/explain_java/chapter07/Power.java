// 计算指数幂

import java.util.Scanner;

class Power {
	// --- 返回x的n次幂 --- //
	strictfp static double power(double x, int n) {
		double tmp = 1.0;
		
		for (int i = 1; i < n; i++)
			tmp *= x;  // 将tmp乘以x

		return tmp;
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("计算a的b次幂。");
		System.out.print("实数a: "); double a = stdIn.nextDouble();
		System.out.print("整数b: "); int b = stdIn.nextInt();

		System.out.println(a + "的" + b + "次幂为" + power(a, b) + "。");
	}
}
