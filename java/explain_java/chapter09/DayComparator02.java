// 日期类Day的使用

import java.util.Scanner;

class DayComparator02 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		int y, m, d;
		System.out.println("请输入日期1。");
		System.out.print("年: "); y = stdIn.nextInt();
		System.out.print("月: "); m = stdIn.nextInt();
		System.out.print("日: "); d = stdIn.nextInt();

		Day day1 = new Day(y, m, d);

		System.out.println("请输入日期2。");
		System.out.print("年: "); y = stdIn.nextInt();
		System.out.print("月: "); m = stdIn.nextInt();
		System.out.print("日: "); d = stdIn.nextInt();

		Day day2 = new Day(y, m, d);

		if (
				day1.getYear() == day2.getYear()
				&& day1.getMonth() == day2.getMonth()
				&& day1.getDate() == day2.getDate()
			)
			System.out.println("相等。");
		else
			System.out.println("不相等。");
			
	}
}
