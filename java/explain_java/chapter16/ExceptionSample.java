// 读入两个整数值，并显示加减乘除运算的结果

import java.util.Scanner;
import java.util.InputMismatchException;

class ExceptionSample {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("对x和y进行加减乘除运算。");

		while (true) {
			try {
				System.out.print("x的值："); int x = stdIn.nextInt();
				System.out.print("y的值："); int y = stdIn.nextInt();

				System.out.println("x + y = " + (x + y));
				System.out.println("x - y = " + (x - y));
				System.out.println("x * y = " + (x * y));
				System.out.println("x / y = " + (x / y));
				System.out.println("x % y = " + (x % y));
			} catch (InputMismatchException e) {
				System.out.println("发生输入错误。" + e);
				String s = stdIn.next();
				System.out.println("忽略了" + s + "。");
			} catch (ArithmeticException e) {
				System.out.println("发生算术错误。 " + e);
				System.out.println("请输入不会发生错误的数值。");
			} finally {
				System.out.println("------------------------");
				System.out.print("再来一次？ (1--Yes/0--No) : ");
				int retry = stdIn.nextInt();
				if (retry == 0) break;
				System.out.println("------------------------");
			}
		}
	}
}
