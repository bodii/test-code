// 图形类群

// -- 图形
abstract class Shape02 {
	abstract void draw();	// 绘图(抽象方法)
	public abstract String toString();
}

// -- 点
class Point02 extends Shape02 {
	Point02() { }

	void draw() {
		System.out.println("+");
	}

	public String toString() {
		return "Point02";
	}
}

// -- 长方形
class Rectangle02 extends Shape02 {
	private int width;
	private int height;

	Rectangle02(int width, int height) {
		this.width = width;
		this.height = height;
	}

	void draw() {
		for (int i = 0; i < height; i++) { 
			for (int j = 0; j < width; j++) 
				System.out.print("*");
			System.out.println();
		}
	}

	public String toString() {
		return String.format(
				"Rectangle02(width: %d, height: %d)", 
				width, 
				height
				);
	}
}

// -- 直线抽象类
abstract class AbstLine extends Shape02 {
	protected int length;

	int getLength() {
		return length;
	}

	void setLength(int length) {
		this.length = length;
	}
}

// -- 横线
class HorzLine extends AbstLine {
	private int length;   // 长度

	HorzLine(int length01) {
		length = length01;
	}

	void draw() {
		for (int i = 1; i <= length; i++) 
			System.out.print("-");
		System.out.println();
	}

	public String toString() {
		return "HorzLine";
	}
}

// -- 竖线
class VertLine extends AbstLine {
	private int length;  // 长度

	VertLine(int length) {
		this.length = length;
	}

	void draw() {
		for (int i = 1; i <= length; i++)
			System.out.println("|");
	}

	public String toString() {
		return "VertLine";
	}
}
