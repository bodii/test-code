// 使用实例方法版来判断闰年

import java.util.Scanner;

class IsLeapTester {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("公历年份: ");
		int y = stdIn.nextInt();
		System.out.println("该年" + (new Day(y).isLeap() ? "是闰年。" : "不是闰年。"));
	}
}
