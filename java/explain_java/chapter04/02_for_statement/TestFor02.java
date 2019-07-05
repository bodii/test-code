// 使用for语句实现将正整数值从0数到正整数值

import java.util.Scanner;

class TestFor02 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("正整数: "); int n = stdIn.nextInt();

		for (int i = 0; i <= n; i++) 
			System.out.println(i);
	}
}
