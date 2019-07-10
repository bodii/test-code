// 确认final形参不可以赋值

class FinalParameter {
	static int sumOf(final int x, final int y, final int r) {
		return x + y + r;
	}

	public static void main(String[] args) {
		System.out.println(sumOf(1, 2, 3));
	}
}
