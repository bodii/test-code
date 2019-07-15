// 汽车类－第2版的使用

import java.util.Scanner;

class CarTester04 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("请输入汽车数据。");
		System.out.print("名称: ");  String name = stdIn.next();
		System.out.print("车宽: ");  int width = stdIn.nextInt();
		System.out.print("车高: ");  int height = stdIn.nextInt();
		System.out.print("车长: ");  int length = stdIn.nextInt();
		System.out.print("燃料数量: ");  double fuel = stdIn.nextDouble();
		System.out.print("购买年份: ");  int y = stdIn.nextInt();
		System.out.print("购买月份: ");  int m = stdIn.nextInt();
		System.out.print("购买日期: ");  int d = stdIn.nextInt();

		Car02 car2 = new Car02(name, width, height, length, fuel, new Day02(y, m, d));

		car2.putSpec();

		System.out.println("购买日期: " + car2.getPurchaseDay());
	}
}
