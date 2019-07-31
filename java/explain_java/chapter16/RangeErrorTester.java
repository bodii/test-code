// 进行1位(0-9)的加法运算

import java.util.Scanner;

// -- 超出范围的异常
class RangeError extends RuntimeException {
	RangeError(int n) { super("超出范围的数值: " + n); }
}

// -- 超出范围的异常(形参)
class ParameterRangeError extends RangeError {
	ParameterRangeError(int n) { super(n); }
}

// 超出范围的异常(返回值)
class ResultRangeError extends RangeError {
	ResultRangeError(int n) { super(n); }
}

public class RangeErrorTester {

	// -- n 为1位(0 - 9)
	static boolean isValid(int n) {
		return n >= 0 && n <= 9;
	}

	// -- 计算1位(0-9)整数a和b的和
	static int add(int a, int b) throws ParameterRangeError, ResultRangeError {
		if (!isValid(a)) throw new ParameterRangeError(a);
		if (!isValid(b)) throw new ParameterRangeError(a);

		int result = a + b;
		if (!isValid(result)) throw new ResultRangeError(result);
		return result;
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("整数a: "); int a = stdIn.nextInt();
		System.out.print("整数b："); int b = stdIn.nextInt();

		try {
			System.out.println("它们的和为" + add(a, b) + "。");
		} catch (ParameterRangeError e) {
			System.out.println("加数超出范围。" + e.getMessage());
		} catch ( ResultRangeError e) {
			System.out.println("计算结果超出范围。" + e.toString());
		}
	}

}
