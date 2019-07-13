class Day {
	private int year;   // 年
	private int month;  // 月
	private int date;   // 日

	// -- 构造函数
	Day(int year, int month, int date) {
		this.year = year; this.month = month;
		this.date = date;
	}

	int getYear() { return year; }

	int getMonth() { return month; }

	int getDate() { return date; }

	void setYear(int year) { this.year = year; }  // 设置年
	void setMonth(int month) { this.month = month; }  // 设置月
	void setDate(int date) { this.date = date; }  // 设置日

	void set(int year, int month, int date) {
		this.year = year;
		this.month = month;
		this.date = date;
	}

	// -- 计算星期
	int dayOfWeek() {
		int y = year;
		int m = month;
		if (m == 1 || m == 2) {
			y--;
			m += 12;
		}

		return (y + y / 4 - y / 100 + y / 400 + (13 * m + 8) / 5 + date) % 7;
	}
}
