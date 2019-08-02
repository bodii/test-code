import java.util.Scanner;

// -- 自已创建的检查异常
class CheckedException extends Exception {
	CheckedException(String s, Throwable e) { super(s, e); }
}

// -- 自已创建的非检查异常
class UncheckedException extends RuntimeException {
	UncheckedException(String s, Throwable e) { super(s, e); }
}

public class Abc {

	// -- 发生sw值所对应的异常
	static void work(int sw) throws Exception {
		switch (sw) {
			case 1: throw new RuntimeException("发生非检查异常!");
			case 2: throw new Exception("发生检查异常!");
		}
	}

	// -- 调用work
	static void test(int sw) throws CheckedException {
		try {
			work(sw);
		} catch (RuntimeException e) {
			throw new UncheckedException("无法处理非检查异常!", e);
		} catch (Exception e) {
			throw new CheckedException("无法处理检查异常!", e);
		}
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("sw: ");
		int sw = stdIn.nextInt();

		try {
			test(sw);
		} catch (Exception e) {
			System.out.println("异常: " + e);
			System.out.println("异常原因: " + e.getCause());
			e.printStackTrace();
		}
	}
}



