// 字符串数组

import java.util.Scanner;

class ReadStringArray {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("字符串的个数: "); int n = stdIn.nextInt();
		String[] sx = new String[n];

		for (int i = 0; i < n; i++) {
			System.out.print("sx[" + i + "] = ");
			sx[i] = stdIn.next();
		}

		for (int i = 0; i < n; i++)
			System.out.println("sx[" + i + "] = \"" + sx[i] + "\"");
	}
}
