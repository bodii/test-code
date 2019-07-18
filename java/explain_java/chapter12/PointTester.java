// 二维坐标类和三维坐标类
//
// 二维坐标类

class Point2D {
	int x;
	int y;

	Point2D(int x, int y) { this.x = x; this.y = y; }

	void setX(int x) { this.x = x; }
	void setY(int y1) { y = y1; }
	int getX() { return x; }
	int getY() { return y; }
}

// 三维坐标类
class Point3D extends Point2D {
	int z;
	
	Point3D(int x1, int y1, int z1) { super(x1, y1); z = z1; }

	void setZ(int z1) { z = z1; }
	int getZ() { return z; }
}

public class PointTester {
	public static void main(String[] args) {
		Point2D a = new Point2D(10, 15);
		Point3D b = new Point3D(20, 30, 40);

		System.out.printf("a = (%d, %d)\n", a.getX(), a.getY());
		System.out.printf("b = (%d, %d, %d)\n" , b.getX(), b.getY(), b.getZ());
	}
}
