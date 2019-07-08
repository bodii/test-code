// 读入数组中全部元素的值并显示

import java.util.Scanner;

class IntArrayScan {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("元素个数: ");
		int n = stdIn.nextInt();
		int[] a = new int[n];

		for (int i = 0; i < n; i++) {
			System.out.printf("a[%d] = ", i); a[i] = stdIn.nextInt();
		}	
		

		for (int i = 0; i < n; i++) 
			System.out.printf("a[%d] = %d\n", i, a[i]);
	}
}
