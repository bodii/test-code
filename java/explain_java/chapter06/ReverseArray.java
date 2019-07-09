// 对数组中的元素进行倒序排列并显示

import java.util.Random;
import java.util.Scanner;

class ReverseArray {
	public static void main(String[] args) {
		Random rand = new Random();
		Scanner stdIn = new Scanner(System.in);

		System.out.print("元素个数: ");
		int n = stdIn.nextInt();
		int[] a = new int[n];
		
		for (int i = 0; i < n; i++) {
			a[i] = 10 + rand.nextInt(90);
			System.out.println("a[" + i + "] = " + a[i]);
		}

		for (int i = 0; i < n / 2; i++) {
			int t = a[i];
			a[i] = a[n - i - 1];
			a[n - i - 1] = t;
		}

		System.out.println("元素的倒序排列执行完毕。");
		for (int i = 0; i < n; i++)
			System.out.println("a[" + i + "] = " + a[i]);
	}
}
