// 计算3个人的身高、体重、年龄的最大值并显示结果

import java.util.Scanner;

class MaxHwaMethod {
	// --- 返回a、b、c中最大值 --- //
	static int max (int a, int b, int c) {
		int max = a;
		if (b > max) max = b;
		if (c > max) max = c;
		return max;
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);
		int[] height = new int[3];
		int[] weight = new int[3];
		int[] age    = new int[3];

		for (int i = 0; i < 3; i++) {
			System.out.print("{" + (i + 1) + "] ");
			System.out.print("身高: ");   height[i] = stdIn.nextInt();
			System.out.print("\t体重: "); weight[i] = stdIn.nextInt();
			System.out.print("\t年龄: "); age[i]    = stdIn.nextInt();
		}

		int maxHeight = max(height[0], height[1], height[2]); // 身高的最大值
		int maxWeight = max(weight[0], weight[1], weight[2]); // 体重的最大值
		int maxAge    = max(age[0], age[1], age[2]);          // 年龄的最大值

		System.out.println("身高的最大值是: " + maxHeight + "。"); 
		System.out.println("体重的最大值是:  " + maxWeight + "。");
		System.out.println("年龄的最大值是:  " + maxAge + "。");

	}

}
