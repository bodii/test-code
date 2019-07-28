// 显示通过命令行参数指定的月份的日历。当只从命令行传入了年份时，显示年从1月份到12月份的日
// 历；当传入了年份和月份时，显示该月的日历；如果年份和月份都未被传入，则显示当前月份的日历。

import java.util.Calendar;

class CalendarTester {
	public static void main(String[] args) {
		int month, year;
		Calendar c = Calendar.getInstance();

		year = !args[0].isEmpty() ?  Integer.parseInt(args[0]) : Calendar.YEAR;
		month = !args[1].isEmpty() ? Integer.parseInt(args[1]) : Calendar.MONTH;
		c.set(year, month);

		System.out.println(c.roll(Calendar.DATE));
	}
}
