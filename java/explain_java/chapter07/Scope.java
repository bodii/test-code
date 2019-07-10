// 确认标识符的作用域
class Scope {
	static int x = 700;  // 类作用域： 字段

	static void printX() {
		System.out.println("x = " + x);
	}

	public static void main(String[] args) {
		System.out.println("x = " + x);

		int x = 800;   // 块作用域: 局部变量
		System.out.println("x = " + x);
		System.out.println("Scope.x = " + Scope.x);
		
		printX();
	}
}
