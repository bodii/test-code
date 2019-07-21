/**
 * 类HorzLine04是表示横线的类。
 * 该类派生自表示直线的抽象类AbstLine。
 * @author wong
 * @see Shape04
 * @see AbstLine04
 */
public class HorzLine04 extends AbstLine04 {

	/**
	 * 创建横线的构造函数。
	 * 接收长度参数。
	 * @param length 创建的直线长度
	 */
	public HorzLine04(int length) { super(length); }

	/**
	 * 访求toString返回表示与横线相关的图形信息的字符串
	 * @return 返回字符串"HorzLine04（length: 3)"。
	 */
	public String toString() {
		return "HorzLine04(length: " + getLength() + ")";
	}

	/**
	 * 方法draw用于绘制横线。
	 * 通过横向排列减号'-'进行绘图。
	 * 连续显示长度个数的'-', 并换行。
	 */
	public void draw() {
		for (int i = 1; i <= getLength(); i++)
			System.out.print('-');
		System.out.println();
	}
}
