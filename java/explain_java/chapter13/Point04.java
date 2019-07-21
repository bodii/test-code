/**
 * 类Point是表示点的类
 * 该类派生自表示图形的抽象类Shape04
 * 无字段
 * @author wong
 * @see Shape04
 */

public class Point04 extends Shape04 {
	/**
	 * 创建点的构造函数。
	 * 不接收参数.
	 */
	public Point04() {
		// 无操作
	}

	/**
	 * 方法toString返回表示与点相关的图形信息的字符串。
	 * 返回的字符串是"Point"
	 * @return 返回字符串"Point"。
	 */
	public String toString() {
		return "Point04";
	}

	/**
	 * 方法draw用于绘制点。
	 * 只显示1个加号'+',并执行
	 */
	public void draw() {
		System.out.println('+');
	}
}
