// 写一段程序，显示读入的数值个符号。* 和 + 交叉显示

import java.util.Scanner;

class TestPutAsterisk01 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("一个正整数: "); int n = stdIn.nextInt();

		while (n > 0) {
			if (n > 0 && (n % 2) == 0) 
				System.out.print('+');
			else
				System.out.print('*');

			n --;
		}

		System.out.println();
	}
}
