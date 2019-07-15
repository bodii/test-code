// 汽车类

public class Car {
	private String name;		// 名称
	private int width;			// 宽度
	private int height;			// 高度
	private int length;			// 长度
	private double x;			// 当前位置的x坐标
	private double y;			// 当前位置的y坐标
	private double fuel;		// 剩余燃料
	private Day02 purchaseDay;	// 购买日期

	// -- 构造函数
	public Car(String name, int width, int height, int length, double fuel,
			Day02 PurchaseDay) {
		this.name = name; this.width = width; this.height = height;
		this.length = length; this.fuel = fuel; x = y = 0.0;
		this.purchaseDay = new Day02(purchaseDay);
	}

	public double getX() { return x; }
	public double getY() { return y; }
	public double getFuel() { return fuel; }
	public Day02 getPurchaseDay() {
		return new Day02(purchaseDay);
	}

	// -- 显示型号
	public void putSpec() {
		System.out.printf("名称: %s\n", name);
		System.out.printf("车宽: %dmm\n", width);
		System.out.printf("车高: %dmm\n", height);
		System.out.printf("车长: %dmm\n", length);
	}

	// -- 向x方向移动dx, 向y方向移动dy
	public boolean move(double dx, double dy) {
		double dist = Math.sqrt(dx * dx + dy * dy);

		if (dist > fuel)
			return false;
		else {
			fuel -= dist;
			x += dx;
			y += dy;
			return true;
		}
	}
}
