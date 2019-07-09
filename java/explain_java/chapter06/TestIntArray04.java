// 使用得连续的值不相等。

import java.util.Random;
import java.util.Scanner;

class TestIntArray04 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);
		Random rand = new Random();

		System.out.print("元素个数: "); int n = stdIn.nextInt();
		int[] a = new int[n];
		for (int i = 0; i < n; i++) {
			a[i] = 1 + rand.nextInt(10);
			if (i > 0 && a[i] == a[i - 1]) {
				do {
					a[i] = 1 + rand.nextInt(10);
				} while (a[i] == a[i - 1]);
			}

			System.out.printf("a[%d] = %d\n", i, a[i]);
		}
		System.out.println();
	}
}
