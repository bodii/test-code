// 显示toString返回的字符串的方法（用于所有的类类型）

class X {
	public String toString() {
		return "Class X";
	}
}

class Y extends X {
	public String toString() {
		return "Class Y";
	}
}

public class ToString {

	// -- 显示toString返回的字符串
	static void print(Object obj) {
		System.out.println(obj);
	}

	public static void main(String[] args) {
		X x = new X();
		Y y = new Y();
		int[] c = new int[5];

		print(x);
		print(y);
		print(c);
	}
}
