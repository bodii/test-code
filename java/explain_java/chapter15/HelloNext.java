// 读入姓名并打招呼

import java.util.Scanner;

class HelloNext {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("姓名: ");
		String s = stdIn.next();

		System.out.println("你好" + s + "先生。");
	}
}
