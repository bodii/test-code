// 线性查找

import java.util.Random;
import java.util.Scanner;

class LinearSearch {
	public static void main(String[] args) {
		Random rand = new Random();
		Scanner stdIn = new Scanner(System.in);

		final int n = 12;
		int[] a = new int[n];

		for (int j = 0; j < n; j++)
			a[j] = rand.nextInt(10);

		System.out.print("数组a中全全部元素的值\n{");
		for (int j = 0; j < n; j++) 
			System.out.printf("%d ", a[j]);
		System.out.println("}");

		System.out.print("要查找的数组: ");
		int key = stdIn.nextInt();
		
		int i;
		for (i = 0; i < n; i++)
			if (a[i] == key)
				break;
		
		if (i < n)
			System.out.println("该元素是a[" + i + "]。");
		else
			System.out.println("该元素不存在。");
	}
}
