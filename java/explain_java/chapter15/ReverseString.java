// 倒序显示该字符串

import java.util.Scanner;

class ReverseString {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("要反转的字符串s1: "); String s1 = stdIn.next();
		
		System.out.print("反转后的字符串s1: ");
		for (int i = s1.length() - 1; i >= 0; i--)
			System.out.print(s1.charAt(i));
		System.out.println();
	}
}
