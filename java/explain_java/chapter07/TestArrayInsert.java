// 编写方法aryIns,将x插入到数组a中的元素a[idx]的位置

import java.util.Scanner;

class TestArrayInsert {

	// -- 在数组中插入元素 -- //
	static void aryIns(int[] a, int idx, int x) {
		for (int i = a.length - 1; i > idx; i--) {
			a[i] = a[i - 1];
		}

		a[idx] = x;
	}

	// -- 打印数组元素 -- //
	static void putArray(int[] a) {
		String s = "数组a = { ";
		for (int i = 0; i < a.length; i++) {
			s += a[i] + ", ";
		}
		s += "}";

		System.out.println(s);
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		int[] a = { 4, 5, 2, 8, 1, 4, 6, 9, 7 };

		putArray(a);

		System.out.print("插入到数组的位置: "); int idx  = stdIn.nextInt();
		System.out.print("插入的元素: "); int x = stdIn.nextInt();

		aryIns(a, idx, x);

		putArray(a);
	}
}
