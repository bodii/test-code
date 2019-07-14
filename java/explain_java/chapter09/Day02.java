// Day类02版

class Day02 {
	private int year;
	private int month;
	private int date;

	public Day02() { set(1, 1, 1); }
	public Day02(int year) { set(year, 1, 1); }
	public Day02(int year, int month, int date) { set(year,month, date); }
	public Day02(Day02 d) { set(d.year, d.month, d.date); }

	// -- 获取年、月、日
	public int getYear() { return year; }
	public int getMonth() { return month; }
	public int getDate() { return date; }

	// -- 设置年、月、日
	public void setYear(int year) { this.year = year; }
	public void setMonth(int month) { this.month = month; }
	public void setDate(int date) { this.date = date; }

	public void set(int year, int month, int date) {
		this.year = year;
		this.month = month;
		this.date = date;
	}

	// -- 计算星期
	public int dayOfWeek() {
		int y = year;
		int m = month;
		if (m == 1 || m == 2) {
			y--;
			m += 12;
		}

		return (y + y / 4 - y / 100 + y / 400 + (13 * m + 8) / 5 + date) % 7;
	}
	
	// -- 与日期d 
	public boolean equalTo(Day02 d) {
		return year == d.year && month == d.month && date == d.date;
	}

	// -- 返回字符串表示
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
