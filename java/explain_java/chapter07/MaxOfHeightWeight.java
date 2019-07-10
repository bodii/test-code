// 计算最高者的身高和最胖者的体重

import java.util.Scanner;

class MaxOfHeightWeight {
	// -- 计算并返回数组a中的最大值 --
	static int maxOf(int[] a) {
		int max = a[0];
		for (int i = 1; i < a.length; i++)
			if (a[i] > max)
				max = a[i];

		return max;

	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("人数: ");
		int ninzu = stdIn.nextInt();

		int[] height = new int[ninzu];
		int[] weight = new int[ninzu];

		System.out.println("请输入" + ninzu + "人的身高和体重。");

		for (int i = 0; i < ninzu; i++) {
			System.out.print((i + 1) + "号的身高: ");
			height[i] = stdIn.nextInt();
			System.out.print((i + 1) + "号的体重: ");
			weight[i] = stdIn.nextInt();
		}

		System.out.println("最高者的身高: " + maxOf(height) + "cm");
		System.out.println("最胖者的体重: " + maxOf(weight) + "kg");
	}
}
