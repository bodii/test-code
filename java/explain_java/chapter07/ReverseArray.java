// 将值读入到数组元素中，并进行倒序排序

import java.util.Scanner;

class ReverseArray {

	// -- 交换数组中的元素a[idx1]和a[idx2] --
	static void swap(int[] a, int idx1, int idx2) {
		int t = a[idx1]; a[idx1] = a[idx2]; a[idx2] = t; 
	}

	// -- 对数组a中的元素进行倒序排列 --
	static void reverse(int[] a) {
		for (int i = 0; i < a.length / 2; i++)
			swap(a, i, a.length - i - 1);
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("元素个数: ");
		int num = stdIn.nextInt();

		int[] x = new int[num];

		for (int i = 0; i < num; i++) {
			System.out.printf("x[%d] : ", i);
			x[i] = stdIn.nextInt();
		}

		reverse(x);

		System.out.println("元素的倒序排列执行完毕。");
		for (int i = 0; i < num; i++)
			System.out.printf("x[%d] = %d\n", i, x[i]);
	}
}
