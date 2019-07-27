// 读入字符串，并显示其全部字符的字符编码

import java.util.Scanner;

class DisplayCodePoint {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("要显示的字符编码的字符串: "); String s = stdIn.next();

		for (int i = 0; i < s.length(); i++)
			System.out.printf(
					"s[%d]: %s 的字符编码是%s\n", 
					i, s.charAt(i), s.getBytes()
					);
	}
}
