// 查找字符串

import java.util.Scanner;

class SearchString {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("字符串s1: "); String s1 = stdIn.next();
		System.out.print("字符串s2: "); String s2 = stdIn.next();

		int idx = s1.indexOf(s2);

		if (idx == -1)
			System.out.println("s1中不包含s2。");
		else
			System.out.printf("s1的第%d个字符串中包含s2。\n", idx + 1);
	}
}
