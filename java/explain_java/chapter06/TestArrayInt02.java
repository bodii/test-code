// 读入数组的元素个数和各个元素的值

import java.util.Scanner;

class TestArrayInt02 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("元素个数: "); int n = stdIn.nextInt();

		int[] a = new int[n];
		String s = "";
		for (int i = 0; i < n; i++) {
			System.out.printf("a[%d] = ", i); a[i] = stdIn.nextInt();
			if (i == n - 1)
				s +=  a[i];
			else
				s +=  a[i] + ", ";

		}

		System.out.println("a = {" + s + "}");

	}
}
