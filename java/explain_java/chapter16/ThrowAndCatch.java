// 用于理解异常处理的示例

import java.util.Scanner;

class ThrowAndCatch {

	// -- 发生sw值所对应的异常
	static void check(int sw) throws Exception {
		switch (sw) {
			case 1: throw new Exception("发生检查异常!!");
			case 2: throw new RuntimeException("发生非检查异常!!");
		}

	}

	// -- 使用check
	static void test(int sw) throws Exception {
		check(sw);
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("sw: ");
		int sw = stdIn.nextInt();

		try {
			test(sw);
		} catch (Exception e) {
			System.out.println(e.getMessage());
		}
	}
}
