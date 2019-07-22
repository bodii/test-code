/**
 * 等腰直角三角形抽象类。
 * 该类派生自表示图形的抽象类Shape04。
 * 具体的直角在左下方、左上方、右下方、右上方的三角形类派生于此类
 * @author wong
 * @see Shape04
 */

public abstract class RightTriangle extends Shape04 {
	
	/**
	 * 直角边
	 */
	private int corner;

	/**
	 * 等腰直角三角形的构造函数。
	 * @param corner 直角边的长度
	 */
	RightTriangle(int corner) {
		setCorner(corner);
	};

	/**
	 * 获取直角边的长度。
	 * @return 直角边的长度
	 */
	public int getCorner() {
		return corner;
	}


	/**
	 * 设置直角边的长度。
	 * @param int 直角边的长度
	 */
	public void setCorner(int corner) {
		this.corner = corner;
	}

	/**
	 * 抽象方法toString返回等腰直角三角形的信息。
	 */
	public abstract String toString(); 
}
