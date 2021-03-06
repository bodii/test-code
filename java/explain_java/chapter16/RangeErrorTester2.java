// 声明异常并捕获

import java.util.Scanner;

// -- 超出最大范围
class RangeError2 extends RuntimeException {
	RangeError2(int n) { super("超出数值的最大的范围：" + n); }
}

// -- 参数超出最大范围
class ParameterRangeError2 extends RangeError2 {
	ParameterRangeError2(int n) { super(n); }
}

// -- 结果超出最大范围
class ResultRangeError2 extends RangeError2 {
	ResultRangeError2(int n) { super(n); }
}

public class RangeErrorTester2 {

	// -- 检测数字是否在0-9
	static boolean isValid (int n) throws ParameterRangeError2 {
		return n >= 0 && n<= 9;
	}

	// -- 合计两个数的值
	static int add(int a, int b) throws ParameterRangeError2, ResultRangeError2 {
		if (!isValid(a)) throw new ParameterRangeError2(a);
		if (!isValid(b)) throw new ParameterRangeError2(b);

		int result = a + b;
		if (!isValid(result)) throw new ResultRangeError2(result);

		return result;
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("整数a: "); int a = stdIn.nextInt();
		System.out.print("整数b: "); int b = stdIn.nextInt();

		try {
			System.out.printf("a = %d, b = %d, a + b = %d\n", a, b, add(a, b));
		} catch (ParameterRangeError2 | ResultRangeError2 e) {
			System.out.println(e.toString());
		}
	}
}
