// -- 测试圆和坐标

public class CircleTester {
	public static void main(String[] args) {
		Point2D[] p = new Point2D[] {
			new Point2D(3, 7), new Point2D(4, 6)
		};
		Circle c1 = new Circle();
		Circle c2 = new Circle(new Point2D(10, 15), 5);

		for (int i = 0; i < p.length; i++) 
			System.out.println("p[" + i + "] = " + p[i].toString());

		c1.setRadius(10);
		System.out.println("c1 = " + c1.toString());
		System.out.println("c2 = " + c2.toString());
	}
}
