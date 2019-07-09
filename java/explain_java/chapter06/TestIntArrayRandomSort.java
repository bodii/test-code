// 将数组中的元素随机排序

import java.util.Random;
import java.util.Scanner;

class TestIntArrayRandomSort {
	public static void main(String[] args) {
		Random rand = new Random();
		Scanner stdIn = new Scanner(System.in);

		System.out.print("元素个数: "); int n = stdIn.nextInt();
		int[] a = new int[n];
		for (int i = 0; i < n; i++) {
			a[i] = 1 + rand.nextInt(10);
			for (int j = 0; j < i; j++) {
				if (a[j] < a[i]) {
					int t = a[i];
					a[i] = a[j];
					a[j] = t;
				}
			}

		}

		for (int i = 0; i < n; i++) {
			System.out.printf("a[%d] = %d\n", i, a[i]);
		}

	}
}
