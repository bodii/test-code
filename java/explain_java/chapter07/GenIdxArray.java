// 创建一个全部元素的值与索引相同的数组

import java.util.Scanner;

class GenIdxArray {

	// -- 创建并返回一个全部元素的值与索引相同、元素个数为n的数组 -- // 
	static int[] idxArray(int n) {
		int[] a = new int[n];
		for (int i = 0; i < n; i++)
			a[i] = i;

		return a;
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("元素个数: "); int n = stdIn.nextInt();
		int[] x = idxArray(n);  // 元素个数为n的数组

		for (int i = 0; i < n; i++)
			System.out.println("x[" + i + "] = " + x[i]);
	}
}
