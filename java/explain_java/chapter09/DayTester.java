// Day类的使用

class DayTester {
	public static void main(String[] args) {
		Day birthday = new Day(1963, 11, 18);

		System.out.printf(
				"生日： %d-%d-%d\n", 
				birthday.getYear(), 
				birthday.getMonth(),
				birthday.getDate()
				);
	}
}
