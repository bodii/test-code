/**
 * 等腰三角形测试类
 *
 * @author wong
 * @see Shape04
 * @see RightTriangle
 */

public class RightTriangleTester {
	public static void main(String[] args) {
		Shape04[] p = new Shape04[4];
		p[0] = new UpperLeftRightTriangle(10);	// 直角在左上角
		p[1] = new RightUpperRightTriangle(10); // 直角在右上角
		p[2] = new LeftLowerRightTriangle(10);	// 直角在左下角
		p[3] = new RightLowerRightTriangle(10); // 直角在右下角

		for (Shape04 s : p) {
			if (s instanceof RightTriangle) {
				System.out.println(s);
				s.draw();
			}
		}
	}
}
