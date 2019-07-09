// 使得元素的值都不相等

import java.util.Scanner;
import java.util.Random;

class TestIntArray05 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);
		Random rand = new Random();
		

		int n;
		do {
			System.out.print("元素个数(大于0小于等于10): "); n = stdIn.nextInt();
		} while (n < 1 || n > 10);
		
		int[] a = new int[n];
		for (int i = 0; i < n; i++) {
			a[i] = 1 + rand.nextInt(10);

			while (i > 0 && in_array(a, a[i], i)) {
				a[i] = 1 + rand.nextInt(10);
			}
			System.out.printf("a[%2$d] = %1$d\n", a[i], i); 
		}
		System.out.println();

	}

	public static boolean in_array(int[] a, int v, int key) {
		for (int i = 0; i < a.length - 1; i++) {
			// System.out.printf("a[%1$d] = %2$d, v=%3$d\n", i, a[i], v);
			if (a[i] == v && i != key) 
				return true;
			else
				continue;
		}

		return false;
	}
}
