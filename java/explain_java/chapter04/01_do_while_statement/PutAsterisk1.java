// 显示所读入的数值个*

import java.util.Scanner;

class PutAsterisk1 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("要显示多少个*呢: ");
		int n = stdIn.nextInt();

		int i = 0;
		while (i < n) {
			System.out.print('*');
			i++;
		}
		System.out.println();
	}
}
