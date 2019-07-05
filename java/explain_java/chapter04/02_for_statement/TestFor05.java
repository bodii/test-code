// 显示读入的数值个*,每显示5个就换行

import java.util.Scanner;

class TestFor05 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("正整数: "); int n = stdIn.nextInt();
		for (int i = 1; i <= n; i++) { 
			System.out.print('*');
			if ((i % 5) == 0)
				System.out.println();
		}

		System.out.println();
	}
}
