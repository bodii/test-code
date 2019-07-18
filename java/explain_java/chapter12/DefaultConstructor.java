// 超类和子类(确认默认构造函数的动作)

// 超类
class A {
	private int a;

	A() { a = 50; }
	int getA() { return a; }
}


// 子类
class B extends A {
	// 未定义构造函数(创建默认构造函数)
}

public class DefaultConstructor {
	public static void main(String[] args) {
		B x = new B();
		
		System.out.println("x.getA() = " + x.getA());
	}
}
