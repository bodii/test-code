// 汽车类的使用

import java.util.Scanner;

class CarTester2 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("请输入车的数据。");
		System.out.print("名称: ");			String name = stdIn.next();
		System.out.print("宽度: ");			int width = stdIn.nextInt();
		System.out.print("高度: ");			int height = stdIn.nextInt();
		System.out.print("长度: ");			int length = stdIn.nextInt();
		System.out.print("燃料数量: ");		double fuel = stdIn.nextDouble();

		Car myCar = new Car(name, width, height, length, fuel);

		while (true) {
			System.out.println("当前位置(" + myCar.getX() + ", " + myCar.getY() +
					") 剩余燃料 " + myCar.getFuel());
			System.out.print("是否移动[0--No/1--Yes]: ");
			if (stdIn.nextInt() == 0) break;

			System.out.print("X方向的移动距离: "); double  dx = stdIn.nextDouble();
			System.out.print("Y方向的移动距离: "); double  dy = stdIn.nextDouble();

			if (!myCar.move(dx, dy))
				System.out.println("燃料不足! ");
		}

	}
}
