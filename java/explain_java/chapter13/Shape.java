// 图形类群

// -- 图形
abstract class Shape {
	abstract void draw();	// 绘图(抽象方法)
}

// -- 点
class Point extends Shape {
	Point() { }

	void draw() {
		System.out.println("+");
	}
}

// -- 长方形
class Rectangle extends Shape {
	private int width;
	private int height;

	Rectangle(int width, int height) {
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
}
