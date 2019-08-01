// 将值读入到数组元素中，并进行倒序排列(存在Bug reverse 再次抛出异常)

import java.util.Scanner;

class ReverseArray5 {

	// -- 交换数级中的元素a[idx1] 和a[idx2]
	static void swap(int[] a, int idx1, int idx2) {
		int t = a[idx1];
		a[idx1] = a[idx2];
		a[idx2] = t;
	}
	
	// -- 对数组a的元素进行倒序排列
	static void reverse(int[] a) {
		try {
			for (int i = 0; i < a.length / 2; i++)
				swap(a, i, a.length - i);
		} catch (ArrayIndexOutOfBoundsException e) {
			throw new RuntimeException("reverse的Bug? ", e);
		}
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("元素个数: "); int n = stdIn.nextInt();

		int[] a = new int[n];

		for (int i = 0; i < n; i++) {
			System.out.printf("a[%d] = ", i); a[i] = stdIn.nextInt();
		}

		try {
			reverse(a);
			System.out.println("元素的倒序排列执行完毕。");
			for (int i = 0; i < n; i++)
				System.out.printf("a[%d] = %d\n", i, a[i]);
		} catch (RuntimeException e) {
			System.out.println("异常 : " + e);
			System.out.println("异常原因：" + e.getCause());
		}
	}

}
