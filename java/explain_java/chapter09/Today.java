// 显示今天的日期

import java.util.GregorianCalendar;
import static java.util.GregorianCalendar.*;

class Today {
	public static void main(String[] args) {
		GregorianCalendar today = new GregorianCalendar();
		System.out.printf(
				"今天是%04d年%02d月%02d日。\n",
				today.get(YEAR), 
				today.get(MONTH) + 1,
				today.get(DATE)
				);

	}
}
