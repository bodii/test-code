// 类/实例 字段和类/实例 方法

class Static {
	private static int s;			// 静态字段(类变量)
	private int a;					// 非静态字段(实例变量)

	public static void m1() {}		// 静态方法(类方法)
	public		  void f1() {}		// 非静态方法(实例方法)

	// -- 静态方法
	public static void m2(int x) {
		s = x;		// 可以访问静态字段
		// a = x;	// 不可以访问非静态字段
		m1();		// 可以调用静态方法
		// f1();	// 不可调用非静态方法
	}

	//  -- 非静态方法
	public void f2(int x) {
		s = x;		// 可以访问静态字段
		a = x;		// 可以访问非静态字段
		m1();		// 可以调用静态方法
		f1();		// 可以调用非静态方法
		System.out.println("s = " + s + "  a = " + a); 
	}
}

public class StaticTester {
	public static void main(String[] args) {
		Static c1 = new Static();
		Static c2 = new Static();

		Static.m2(5);
		c1.f2(1);
		c2.f2(20);
	}
}
