// 将随机数赋给数组中的全部元素，并显示为横向柱形图

import java.util.Random;
import java.util.Scanner;

class IntArrayRand {
	public static void main(String[] args) {
		Random rand = new Random();
		Scanner stdIn = new Scanner(System.in);

		System.out.print("元素个数: "); int n = stdIn.nextInt();
		int[] a = new int[n];

		for (int i = 0; i < n; i++) 
			a[i] = 1 + rand.nextInt(10); // 1-10 的随机数

		for (int i = 0; i < n; i++) {
			System.out.printf("a[%d] : ", i);
			for (int j = 0; j < a[i]; j++)
				System.out.print('*');
			System.out.println();
		}
	}
}
