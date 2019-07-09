// 计算double 型数组中全部元素的合计值和平均值。元素个数和全部元素的值都通过键盘输入。

import java.util.Scanner;

strictfp class TestDoubleArray03 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);
		
		System.out.print("元素个数:"); int n = stdIn.nextInt();

		double[] a = new double[n];
		double sum = 0.0;
		for (int i = 0; i < n; i++) {
			System.out.printf("a[%d] = ", i); a[i] = stdIn.nextDouble();
			sum += a[i];
		}

		System.out.println("合计值:" + sum);
		System.out.println("平均值:" + (sum / 2));
	}
}
