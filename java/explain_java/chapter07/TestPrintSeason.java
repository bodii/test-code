// 显示参数m所指定的月份的季节。如果m为3、4、5,则显示"春天"...如果为其他值，则不显示任何内容.

import java.util.Scanner;

class TestPrintSeason {
	static void printSeason(int m) {
		if (3 <= m && m <= 5) 
			System.out.println("春天");
		else if (6 <= m && m <= 8)
			System.out.println("夏天");
		else if (9 <= m && m <= 11)
			System.out.println("秋天");
		else if (12 == m || 1 == m || 2 == m)
			System.out.println("冬天");
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		boolean exit = false;
		while (!exit)  {
			System.out.print("输入月份(0退出): "); int m = stdIn.nextInt();
			if (0 == m) exit = true;
			printSeason(m);
		}
	}
}
