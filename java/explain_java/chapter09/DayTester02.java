// 日期类Day的使用

import java.util.Scanner;

class DayTester02 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);
		String[] wd = {"日", "一", "二", "三", "四", "五", "六" };

		System.out.println("请输入阳历生日。");
		System.out.print("年: "); int y = stdIn.nextInt();
		System.out.print("月: "); int m = stdIn.nextInt();
		System.out.print("日: "); int d = stdIn.nextInt();

		Day birthday = new Day(y, m, d);

		System.out.printf(
				"你的生日 %d年 %d月 %d日是星期%s。\n",
				birthday.getYear(),
				birthday.getMonth(),
				birthday.getDate(),
				wd[birthday.dayOfWeek()]
				);
	}
}
