// 超类和子类

// 超类
class Base {
	protected int x;	// 限制公开(该类和下位类可以访问)

	Base() { this.x = 0; }
	Base(int x) { this.x = x; }

	void print() { System.out.println("Base.x = " + x); }
}

// 子类
class Derived extends Base {
	int x;     // 与超类同名的字段

	Derived(int x1, int x2) { super.x = x1; this.x = x2; }

	// 重写超类的方法
	void print() { super.print(); System.out.println("Derived.x = " + x); }
}


public class SuperTester {
	public static void main(String[] args) {
		Base a = new Base(10);
		System.out.println("-- a -- "); a.print();

		Derived b = new Derived(20, 30);
		System.out.println("-- b --"); b.print();
	}
}
