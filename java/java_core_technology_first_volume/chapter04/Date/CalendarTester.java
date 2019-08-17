import java.util.GregorianCalendar;
import java.util.Calendar;

public class CalendarTester {
	public static void main(String[] args) {
		GregorianCalendar someDay = new GregorianCalendar(2019, 8, 12);
		someDay.add(Calendar.DAY_OF_MONTH, 1000);

		int year = someDay.get(Calendar.YEAR);
		int month = someDay.get(Calendar.MONTH) + 1;
		int day = someDay.get(Calendar.DAY_OF_MONTH);

		System.out.printf("year = %d, month = %d, day = %d\n", year, month, day);
	}
}
