/**
 * RightLowerRightTriangle 类
 * 直角在左下方的三角形类
 * @author wong
 * see Shape04
 * see RightTriangle
 */

public class RightLowerRightTriangle extends RightTriangle {

	/**
	 * 创建右下方的三角形的构造函数。
	 * 接收直角边和斜角边参数。
	 * @param corner 直角边
	 */
	public RightLowerRightTriangle(int corner) {
		super(corner);
	}

	/**
	 * 方法toString返回表示左下方的等腰直角三角形信息
	 * @return 返回字符串RightLowerRightTriangle(corner: 3)"。
	 */
	public String toString() {
		return String.format(
				"RightLowerRightTriangle(corner: %d)",
				getCorner()
				);
	}

	/**
	 * 方法draw用于绘制左下方法的三角形。
	 */
	public void draw() {
		for(int i = 1; i <= getCorner(); i++) {
			for(int j = getCorner() - i; j > 0; j--)
				System.out.print(' ');
			for(int j = 1; j <= i; j++)
				System.out.print('*');
			System.out.println();
		}
	}

}
