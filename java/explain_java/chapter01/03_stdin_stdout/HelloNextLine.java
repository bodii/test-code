// 读入姓名并打招呼

import java.util.Scanner;

class HelloNextLine {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("您的姓名是: ");
		String s = stdIn.nextLine();      // 读入1行字符串

		System.out.println("你好" + s + "先生");
	}
}
