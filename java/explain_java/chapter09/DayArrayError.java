// 日期类Day的数组

import java.util.Scanner;

class DayArrayError {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);
		String[] wd = {"日", "一", "二", "三", "四", "五", "六" };

		System.out.print("日期个数: ");
		int n = stdIn.nextInt();

		Day[] a = new Day[n];    // 元素个数为n的Day类型数组
		for (int i = 0; i < n; i++)
			a[i].set(2017, 10, 15);
			// a[i] = new Day(2017, 10, 15); // ok
			

		for (int i = 0; i < a.length; i++)
			System.out.printf(
					"a[%d] = %d年%d月%d日(星期%s)\n",
					i,
					a[i].getYear(), // err 未实例
					a[i].getMonth(),
					a[i].getDate(),
					a[i].dayOfWeek()
					);
	}
}
