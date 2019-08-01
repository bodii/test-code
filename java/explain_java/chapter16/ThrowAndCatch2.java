// 用于异常处理的练习

import java.util.Scanner;

class ExceptionError extends Exception {
	ExceptionError(String msg) { super(msg); }
}

class RuntimeError extends RuntimeException {
	RuntimeError(String msg) { super(msg); }
}


class ThrowAndCatch2 {

	// -- 发生sw值的对应异常
	static void check(int sw) throws ExceptionError, RuntimeError {
		switch (sw) {
			case 1: throw new ExceptionError("发生检查异常!");
			case 2: throw new RuntimeError("发生非检查异常!!");
		}
	}

	static void test(int sw) throws ExceptionError, RuntimeError {
		check(sw);
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("sw： "); int sw = stdIn.nextInt();

		try {
			test(sw);
		} catch (ExceptionError e) {
			System.out.println(e.toString());
		} catch (RuntimeError e) {
			System.out.println(e.toString());
		} finally {
			System.out.printf("sw = %d\n", sw);
		}
	}
}
