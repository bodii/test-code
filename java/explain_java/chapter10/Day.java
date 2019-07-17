// 日期类Day

public class Day {
	private int year = 1;
	private int month = 1;
	private int date = 1;

	// -- y年是闰年吗
	public static boolean isLeap(int y) {
		return y % 4 == 0 && y % 100 != 0 || y % 400 == 0;
	}

	// -- 构造函数
	public Day()							{}
	public Day(int year)					{ this.year = year; }
	public Day(int year, int month)			{ this(year); this.month = month;  }
	public Day(int year, int month, int date) {this(year, month); this.date = date; }
	public Day(Day d)						{ this(d.year, d.month, d.date); }

	// -- 获取年、月、日
	public int getYear()	{ return year; }
	public int getMonth()   { return month; }
	public int getDay()		{ return date; }

	// -- 设置年、月、日
	public void setYear(int year)   { this.year = year; }
	public void setMonth(int month) { this.month = month; }
	public void setDate(int date)	{ this.date = date; }

	public void set(int year, int month, int date) {
		this.year = year;
		this.month = month;
		this.date = date;
	}

	// -- 是闰年吗
	public boolean isLeap() { return isLeap(year); }

	// -- 计算星期
	public int dayOfWeek() {
		int y = year;
		int m = month;
		if (m == 1 || m == 2) {
			y--;
			m += 12;
		}

		return (y + y / 4 - y / 100 + y / 400 + ( 13 * m + 8) / 5 + date) % 7;
	}

	// -- 与日期d相等吗
	public boolean equalTo(Day d) {
		return year == d.year && month == d.month && date == d.date;
	}

	// -- 返回字符串的表示
	public String toString() {
		String[] wd = { "日", "一", "二", "三", "四", "五", "六" };
		return String.format(
				"%04d年%02d月%02d日(%s)",
				year,
				month,
				date,
				wd[dayOfWeek()]
				);

	}
}
