/**
 * 类Rectangle04类是表示长方形的类。
 * 该类派生自表示图形的抽象类Shape04。
 * @author wong
 * @see Shape04
 */
public class Rectangle04 extends Shape04 {
	/**
	 * 表示长方形的int型字段。
	 */
	private int width;

	/**
	 * 表示长方形的宽的int型字段。
	 */
	private int height;

	/**
	 * 创建长方形的构造函数。
	 * 接收长和宽作为参数。
	 * @param width 长方形的长
	 * @param height 长方形的宽。
	 */
	public Rectangle04(int width, int height) {
		this.width = width;
		this.height = height;
	}

	/**
	 * 方法toString返回表示与长方形相关的图形信息的字符串
	 * @return 返回字符串Rectangle04(width: 4, height: 3)"。
	 */
	public String toString() {
		return String.format("VerLine04(width: %d, height: %d", width, height);
	}

	/**
	 * 方法draw用于绘制长方形。
	 * 通过拜请星号'*'进行绘图。
	 * 循环width次显示长度个数的'*'并换行。
	 */
	public void draw() {
		for (int i = 1; i <= height; i++) {
			for (int j = 1; j <= width; j++) {
				System.out.print('*');
			}
			System.out.println();
		}
	}

}
