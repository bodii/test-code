// 日期类Day的数组

class DayArrayInit {
	public static void main(String[] args) {
		String[] wd = {"日", "一", "二", "三", "四", "五", "六"};

		// 明治、大正、昭和、平成的开始日期
		Day[] x = {
			new Day(1868, 9, 8),
			new Day(1912, 7, 30),
			new Day(1926, 12, 25),
			new Day(1989, 1, 8),
		};

		for (int i = 0; i < x.length; i++)
			System.out.printf(
					"x[%d] = %d年%d月%d日(星期%s)\n",
					i,
					x[i].getYear(),
					x[i].getMonth(),
					x[i].getDate(),
					x[i].dayOfWeek()
					);
	}
}
