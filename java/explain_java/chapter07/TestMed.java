// 编写方法med,计算3个int型参数a、b、c中的中间值

class TestMed {
	// -- 返回a、b、c的中间值
	static int med (int a, int b, int c) {
		if ((a - b) * (b - c) > 0) return b;		// a大于b && b大于c
		else if ((b - a) * (a - c) > 0) return a;	// b大于a && a大于c 
		else return c;
	}

	public static void main(String[] args) {
		int a = 10, b = 3, c = 6;

		System.out.printf("a = %d, b = %d, c = %d, 中间值是: %d\n", a, b, c, med(a, b, c));
	}
}
