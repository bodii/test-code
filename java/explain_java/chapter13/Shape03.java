// 图形类群

// -- 图形
abstract class Shape03 {
	abstract void draw();	// 绘图(抽象方法)
	public abstract String toString();

	void print() {
		System.out.println(toString());
		draw();
	}
}

// -- 点
class Point03 extends Shape03 {
	Point03() { }

	void draw() {
		System.out.println("+");
	}

	public String toString() {
		return "Point02";
	}
}

// -- 长方形
class Rectangle03 extends Shape03 {
	private int width;
	private int height;

	Rectangle03(int width, int height) {
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
				"Rectangle03(width: %d, height: %d)", 
				width, 
				height
				);
	}
}

// -- 直线抽象类
abstract class AbstLine02 extends Shape03 {
	private int length;

	AbstLine02(int length) {
		setLength(length);
	}

	int getLength() {
		return length;
	}

	void setLength(int length) {
		this.length = length;
	}
}

// -- 横线
class HorzLine02 extends AbstLine02 {

	HorzLine02(int length01) {
		super(length01);
	}

	void draw() {
		for (int i = 1; i <= getLength(); i++) 
			System.out.print("-");
		System.out.println();
	}

	public String toString() {
		return "HorzLine02";
	}
}

// -- 竖线
class VertLine02 extends AbstLine02 {

	VertLine02(int length) {
		super(length);
	}

	void draw() {
		for (int i = 1; i <= getLength(); i++)
			System.out.println("|");
	}

	public String toString() {
		return "VertLine";
	}
}
