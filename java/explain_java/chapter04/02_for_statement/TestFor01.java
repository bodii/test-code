// 用for语句来实现将正整数值倒数到0的代码

import java.util.Scanner;

class TestFor01 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("正整数值: "); int n = stdIn.nextInt();
		for (int i = n; i >= 0; i--) 
			System.out.println(i);
	}
}
