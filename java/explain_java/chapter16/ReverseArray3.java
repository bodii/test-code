// 将值读入到数组元素中，并进行倒序倒序排列(存在Bug, 在reverse中捕获异常)

import java.util.Scanner;

class ReverseArray3 {

	// -- 交换数组中的元素a[idx1]和a[idx2]
	static void swap(int[] a, int idx1, int idx2) {
		int t = a[idx1];
		a[idx1] = a[idx2];
		a[idx2] = t;
	}

	// -- 对数组a的元素时行倒序排列
	static void reverse(int[] a) {
		try {
			for (int i = 0; i < a.length / 2; i++)
				swap(a, i, a.length - i); // a.length - i - 1
		} catch (ArrayIndexOutOfBoundsException e) {
			e.printStackTrace();
			System.exit(1);
		}
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("元素个数："); int n = stdIn.nextInt();

		int[] arr = new int[n];
		for (int i = 0; i < n; i++) {
			System.out.printf("a[%d] = ", i); arr[i] = stdIn.nextInt();
		}

		reverse(arr); // 排序

		for (int i = 0; i < n; i++)
			System.out.printf("a[%d] = %d\n", i, arr[i]);

	}
}
