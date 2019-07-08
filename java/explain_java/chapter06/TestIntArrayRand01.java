// 显示纵向柱形图
// 最下面一行显示索引除以10的余数。

import java.util.Random;
import java.util.Scanner;

class TestIntArrayRand01 {
	public static void main(String[] args) {
		Random rand = new Random();
		Scanner stdIn = new Scanner(System.in);

		System.out.print("元素个数: "); int n = stdIn.nextInt();
		int[] a = new int[n];

		for (int i = 0; i < n; i++) 
			a[i] = 1 + rand.nextInt(10);

		for (int j = 12; j > 0; j--) {
			for (int i = 0; i < n; i++) {
				if ((j - 2) <= a[i] && j > 2)
					System.out.printf("%-3s", '*');
				else if ((j - 2) > a[i]) 
					System.out.printf("%3s", " ");
				else if (j == 2)
					System.out.printf("%3s", "---");
				else if (j == 1)
					System.out.printf("%1d  ", i % 10);
			}
			System.out.println();
		}

	}
}
