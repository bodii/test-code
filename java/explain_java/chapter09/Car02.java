// 汽车类－第2版

public class Car02 {
	private String name;
	private int width;
	private int height;
	private int length;
	private double x;
	private double y;
	private double fuel;		// 剩余燃料
	private Day02 purchaseDay;  // 购买日期

	// -- 构造函数
	public Car02(String name, int width, int height, int length, double fuel, 
			Day02 purchaseDay) {
		this.name = name;   this.width = width;		this.height = height;
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
		System.out.println("名称: " + name);
		System.out.println("车宽: " + width + "mm");
		System.out.println("车高: " + height + "mm");
		System.out.println("车长: " + length + "mm");
	}

	// -- 向x方向移动dx、向y方向移动dy
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
