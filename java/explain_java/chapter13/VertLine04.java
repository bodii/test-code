/**
 * 类VertLine是表示竖线的类
 * 该类派生自表示直线的抽象类AbstLine。
 * @author wong
 * @see Shape04
 * @see AbstLine04
 */
public class VertLine04 extends AbstLine04 {

	/**
	 * 创建竖线的构造函数。
	 * 接收长度参数。
	 * @param length 创建的直线长度。
	 */
	public VertLine04(int length) { super(length); }

	/**
	 * 方法toString返回表示与竖线相关的图形信息的字符串。
	 * @return 返回字符串“VertLine(length: 3)"。
	 */
	public String toString() {
		return "VertLine04(length: " + getLength() + ")";
	}

	/**
	 * 方法draw用于绘制竖线。
	 * 通过纵向排列竖线'|‘进行绘图。
	 * 循环显示长度个数的‘|‘及换行。
	 */
	public void draw() {
		for (int i = 1; i <= getLength(); i++)
			System.out.println('|');
	}	

}
