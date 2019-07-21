/**
 * 类Shape04是表示图形概念的抽象类
 * 由于为抽象类，因此无法创建该类的实例。
 * 具体的图形类从该类进行派生。
 * @author wong
 * @see Object
 */

public abstract class Shape04 {

	/**
	 * 返回表示图形信息的字符串的抽象方法。
	 * 在类Shape派生的类中实现该方法的主体。
	 * 该方法将java.lang.Object类中的方法重写为抽象方法。
	 */
	public abstract String toString();

	/**
	 * 方法draw是用来绘制图形的抽象方法。
	 * 在类Shape派生的类中实现该方法的主体。
	 */
	public abstract void draw();

	/**
	 * 方法print用来显示图形信息，并绘制图形。
	 * 具体分为如下两步，按顺序依次执行。
	 * Step 1. 显示方法toString返回的字符串，并执行。
	 * Step 2. 调用方法draw来绘制图形。
	 */
	public void print() {
		System.out.println(toString());
		draw();
	}
}
