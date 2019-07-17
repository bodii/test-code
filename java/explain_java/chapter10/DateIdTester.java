// 标识编号类

import java.util.GregorianCalendar;
import static java.util.GregorianCalendar.*;

class DateId {
	private static int counter;	// 赋到了哪一个标识编号
	private int id;		// 标识编号

	static {
		GregorianCalendar today = new GregorianCalendar();
		int y = today.get(YEAR);	// 年
		int m = today.get(MONTH) + 1;	// 月
		int d = today.get(DATE);	// 日

		System.out.printf("今天是%04d年%02d月%02d日。\n", y, m, d);
		counter = y * 1000000 + m * 10000 + d * 100;
	}

	// -- 构造函数
	public DateId() {
		id = ++counter;
	}

	// -- 获取标识编号
	public int getId() {
		return id;
	}
}


public class DateIdTester {
	public static void main(String[] args) {
		DateId a = new DateId();
		DateId b = new DateId();
		DateId c = new DateId();

		System.out.println("a的标识编号: " + a.getId());
		System.out.println("b的标识编号: " + b.getId());
		System.out.println("c的标识编号: " + c.getId());
	}
}
