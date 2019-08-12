import java.time.LocalDate;

public class LocalDateTester {
	public static void main(String[] args) {
		LocalDate now = LocalDate.now();

		// 设置一个日期
		LocalDate newYearsEve = LocalDate.of(2019, 8, 12);
		int year = newYearsEve.getYear();
		int month = newYearsEve.getMonthValue();
		int day = newYearsEve.getDayOfMonth();

		System.out.printf("year = %d, month = %d, day = %d\n", year, month, day);

		// 在当前日期的基础上加1000天
		LocalDate aThousandDaysLater = newYearsEve.plusDays(1000);
		year = aThousandDaysLater.getYear();
		month = aThousandDaysLater.getMonthValue(); 
		day = aThousandDaysLater.getDayOfMonth();

		System.out.printf("year = %d, month = %d, day = %d\n", year, month, day);

	}
}

