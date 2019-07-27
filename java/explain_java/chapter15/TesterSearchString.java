// 查找字符串，相同的部分要上下对齐显示

import java.util.Scanner;

class TesterSearchString {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("字符串s1: "); String s1 = stdIn.next();
		System.out.print("字符串s2: "); String s2 = stdIn.next();

		int in = s1.indexOf(s2);

		if (in < 0) 
			System.out.println("s1不包含s2");
		else {
			System.out.println(s1);
			System.out.printf(
					"%s%s%s。\n", 
					" ".repeat(in), s2, 
					" ".repeat(s1.length() - (in + s2.length()))
					);
		}
	}
}
