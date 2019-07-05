// 程序会逆序显示读入的正整数值，例如，当输入1254时，显示4521

import java.util.Scanner;

class TestReverseNumber {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("输入一个正整数: "); int n = stdIn.nextInt();
		String s = n + "";
		int l = s.length();
		while (l > 0) {
			l --;
			System.out.print(s.charAt(l));
		}

		System.out.println();
	}
}
