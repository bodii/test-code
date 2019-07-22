/**
 * UpperLeftRightTriangle 类
 * 直角在左上方的三角形类
 * @author wong
 * see Shape04
 * see RightTriangle
 */

public class UpperLeftRightTriangle extends RightTriangle {

	/**
	 * 创建左下方的三角形的构造函数。
	 * 接收直角边和斜角边参数。
	 * @param corner 直角边
	 */
	public UpperLeftRightTriangle(int corner) {
		super(corner);
	}

	/**
	 * 方法toString返回表示左下方的等腰直角三角形信息
	 * @return 返回字符串UpperLeftRightTriangle(corner: 3)"。
	 */
	public String toString() {
		return String.format(
				"UpperLeftRightTriangle(corner: %d)",
				getCorner()
				);
	}

	/**
	 * 方法draw用于绘制左下方法的三角形。
	 */
	public void draw() {
		for(int i = getCorner(); i > 0; i--) {
			for(int j = 1; j <= i; j++)
				System.out.print('*');
			System.out.println();
		}
	}

}
