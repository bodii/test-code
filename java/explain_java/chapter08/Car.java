// 汽车类

class Car {
	private String name; 
	private int width;
	private int height;
	private int length;
	private double x;		// 当前位置的x坐标
	private double y;		// 当前位置的y坐标
	private double fuel;	// 剩余燃料

	// 构造函数 
	Car(String name, int width, int height, int length, double fuel) {
		this.name = name;	this.width = width;		this.height = height;
		this.length = length;	this.fuel = fuel;
		x = y = 0.0;
	}

	double getX() { return x; }
	double getY() { return y; }
	double getFuel() { return fuel; }

	// 显示型号
	void putSpec() {
		System.out.println("名称: " + name);
		System.out.println("车宽: " + width + "mm");
		System.out.println("车高: " + height + "mm");
		System.out.println("车长: " + length + "mm");
	}

	// 向x方向移动dx、向y方向移动dy
	boolean move(double dx, double dy) {
		double dist = Math.sqrt(dx * dx +  dy * dy);

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
