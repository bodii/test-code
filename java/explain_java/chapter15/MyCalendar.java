// 自定义日历输出类

import java.util.Calendar;

public class MyCalendar {
	/**
	 * 获取某月中的某天是星期几
	 * 
	 * @param Calendar cal 天数
	 * @return 某月中的某天是星期几
	 */
	public static int weekDay(Calendar cal) {
		int weekDay = cal.get(Calendar.DAY_OF_WEEK);
		if (weekDay == 1)  // 西方星期日为第一天，星期一为第二天
			weekDay = 7;
		else
			weekDay -= 1;
		return weekDay;
	}

	/** 
	 * 打印某年的所有日历
	 * @param int year 某年
	 */
	public void yearCalendar(int year) {
		for (int i = 1; i <= 12; i++) {
			calendar(year, i);
			System.out.println("\n");
		}
	}

	/**
	 * 打印某年某月的日历
	 *
	 * @param int year 某年
	 * @param int month 某月
	 *
	 */
	public void calendar(int year, int month) {
		Calendar firstCal = Calendar.getInstance();
		Calendar lastCal = Calendar.getInstance();

		System.out.println("\t\t" + year + "年" + month + "月");
		System.out.println("日\t一\t二\t三\t四\t五\t六");

		firstCal.set(year, month - 1,  1);
		int dateOfMonth = firstCal.getActualMaximum(Calendar.DATE); // 获取该月的天数
		lastCal.set(year, month - 1, dateOfMonth); // 获取该月的最后一天

		// 获取所求月第一天是星期几，输出是中文而不是数字的
		int weekOfMonth = firstCal.getActualMaximum(Calendar.WEEK_OF_MONTH); // 获取该月的星期数
		String[][] week = new String[weekOfMonth][7];
		int firstDay = weekDay(firstCal); // 获取所求月第一天是星期几
		int lastDay = weekDay(lastCal);   // 获取所求月最后一天是星期几

		int m = 1,  f = 1; // f的作用主要是判断是否需要将第一个星期归到中间几个星期一起计算
		int j;  // 第一个星期
		if (firstDay == 7)
		{
			f = 0;
		} else {
			for (j = 0; j < firstDay; j++) {
				week[0][j] = " ";
				System.out.print(week[0][j] + "\t");
			}

			for (j = firstDay; j < 7; j++) {
				week[0][j] = m + "";
				m++;
				System.out.print(week[0][j] + "\t");
			}

			System.out.println();
		}

		// 中间的几个星期
		for (int i = f; i < weekOfMonth - 1; i++) {
			for (j = 0; j < 7; j++) {
				week[i][j] = m + "";
				m++;
				System.out.print(week[i][j] + "\t");
			}
			System.out.println();
		}

		// 最后一个星期
		if (lastDay == 7) {
			week[weekOfMonth - 1][0] = m + "";
			System.out.print(week[weekOfMonth - 1][0] + "\t");
		} else {
			for (j = 0; j <= lastDay; j++) {
				week[weekOfMonth-1][j] = m + "";
				m++;
				System.out.print(week[weekOfMonth - 1][j] + "\t");
			}
		}

	}
}
