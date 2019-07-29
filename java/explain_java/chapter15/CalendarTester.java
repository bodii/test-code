// 显示通过命令行参数指定的月份的日历。当只从命令行传入了年份时，显示年从1月份到12月份的日
// 历；当传入了年份和月份时，显示该月的日历；如果年份和月份都未被传入，则显示当前月份的日历。

import java.util.Calendar;

class CalendarTester {
	public static void main(String[] args) {
		int month = 0, year = 0;
		Calendar c = Calendar.getInstance();
		MyCalendar mcal = new MyCalendar();
		
		System.out.println("<<<<<<<<<<万年历>>>>>>>>>");
		if (args.length >= 2) {
			year = Integer.parseInt(args[0]);
			month = Integer.parseInt(args[1]);
			mcal.calendar(year, month);
		} else if (args.length == 1) {
			year = Integer.parseInt(args[0]);
			mcal.yearCalendar(year);
		} else if (args.length == 0) {
			year =  c.get(Calendar.YEAR);
			month = c.get(Calendar.MONTH) + 1;
			mcal.calendar(year, month);
		}

		System.out.println();
	}
}
