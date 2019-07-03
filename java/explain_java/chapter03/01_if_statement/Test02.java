// 读入一个正整数值，如果它可以被5整除，则显示“该值可以被5整除“，否则显示"该值不可以被5整除"

import java.util.Scanner;

class Test02 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("输入值: "); int i = stdIn.nextInt();
		if ((i % 5) == 0)
			System.out.println("该值可以被5整除");
		else
			System.out.println("该值不可以被5整除");

	}
}
