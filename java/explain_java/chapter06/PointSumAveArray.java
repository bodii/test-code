// 读入分数并显示总分和平均分(使用数组的版本)

import java.util.Scanner;

class PointSumAveArray {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);
		int sum = 0;           // 总分
		final int ninzu = 5;   // 人数
		int[] tensu = new int[ninzu]; // 分数

		System.out.println("请输入" + ninzu + "名学生的分数。");
		for ( int i = 0; i < ninzu; i++) {
			System.out.print((i + 1) + "号的分数: ");
			tensu[i] = stdIn.nextInt();
			sum += tensu[i];
		}

		System.out.println("总分为" + sum + "分。");
		System.out.println("平均分为" + (double) (sum / ninzu) + "分。");
	}
}
