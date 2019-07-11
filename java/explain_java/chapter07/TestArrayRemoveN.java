// 编写方法aryRmvN,删除数组a中从元素a[idx]开始的n个元素

import java.util.Scanner;

class TestArrayRemoveN {

	// -- 删除数组元素 -- //
	static void aryRmvN(int[] a, int idx, int n) {
		for (int i = idx; i < (idx + n); i++)
			a[i] = a[idx + n];
	}

	// -- 打印数组元素 -- //
	static void putArray(int[] a) {
		String s = "a = { ";
		for(int i = 0; i < a.length; i++) {
			s += a[i] + ", ";
		}
		s += "}";
		System.out.println(s);
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		int[] a = { 3, 7, 2, 9, 4, 6, 8 };

		putArray(a);	

		System.out.printf("删除元素a开始于: "); int start = stdIn.nextInt();
		System.out.printf("删除几个元素: "); int n = stdIn.nextInt();
		aryRmvN(a, start - 1, n);
		
		putArray(a);

	}
}
