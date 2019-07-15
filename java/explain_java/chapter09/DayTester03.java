// 日期类Day

import java.util.Scanner;

class DayTester03 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("请输入day1。");
		System.out.print("年: "); int y = stdIn.nextInt();
		System.out.print("月: "); int m = stdIn.nextInt();
		System.out.print("日: "); int d = stdIn.nextInt();

		Day02 day1 = new Day02(y, m, d);
		System.out.println("day1 = " + day1);

		Day02 day2 = new Day02(day1);
		System.out.println("创建了与day1的日期相同的day2。");

		if (day1.equalTo(day2))
			System.out.println("day1和day2相等。");
		else
			System.out.println("day1和day2不相等。");

		Day02 d1 = new Day02();
		Day02 d2 = new Day02(2010);
		Day02 d3 = new Day02(2010, 10);
		Day02 d4 = new Day02(2010, 10, 15);

		System.out.println("d1  = " + d1);
		System.out.println("d2  = " + d2);
		System.out.println("d3  = " + d3);
		System.out.println("d4  = " + d4);

		Day02[] a = new Day02[3];
		for (int i = 0; i < a.length; i++)
			a[i] = new Day02();

		for (int i = 0; i < a.length; i++)
			System.out.printf("a[%d] = %s\n", i , a[i]);

	}
}
