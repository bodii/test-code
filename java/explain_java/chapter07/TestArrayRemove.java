// 编写方法aryRmv,删除数组a中的元素a[idx]

import java.util.Scanner;

class TestArrayRemove {

	// -- 将数组元素，从idx开始向前移一位
	static void leftMove(int[] a, int idx) {
		a[idx] = a[idx + 1]; 
	}

	// -- 删除数组a中的元素a[idx] 
	static void aryRmv(int[] a, int idx) {
		for (int i = idx; i < a.length - 1; i++) {
			leftMove(a, i);
		}
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		int[] a = { 1, 7, 4, 2, 6, 9, 5, 3};

		String s = "a = { ";
		for (int i = 0; i < a.length; i++)
			s += a[i] + ", ";
		s += " }";
		System.out.println("数组: " + s);

		System.out.print("删除第几个元素: ");
		int n = stdIn.nextInt();
		if (n > 0 && n <= a.length)
			aryRmv(a, n - 1);
		else
			System.out.println("超出数组的索引范围。");

		s = "a = { ";
		for (int i = 0; i < a.length; i++)
			s += a[i] + ", ";
		s += " }";

		System.out.println(s);
	}
}
