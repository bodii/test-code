// 显示输入的月份所处的季节

import java.util.Scanner;

class Season {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		int retry;   // 要重复一次吗
		do {
			System.out.print("计算季节。\n请输入月份: ");
			int month = stdIn.nextInt();

			if (month >= 3 && month <= 5)
				System.out.println("这是春天。");
			else if (month >=6 && month <= 8)
				System.out.println("这是夏天。");
			else if (month >=9 && month <= 11)
				System.out.println("这是秋天。");
			else if (month == 12 || month == 1 || month == 2)
				System.out.println("这是冬天。");

			System.out.print("要重复一次吗? 1--Yes / 0--No: ");
			retry = stdIn.nextInt();

		} while (retry == 1);
	}
}
