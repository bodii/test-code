/**
 * 递归算法
 */

public class Recursive {

	public static int factoria(int n) {
		int value = 1;
		while (n > 0) {
			value *= n--;
		}

		return value;
	}

	public static int recur(int n) {
		if (n == 0)
			return 1;
		else
			return n * recur(n - 1);
	}

	public static void main(String[] args) {
		// 测试
		int v1 = factoria(10);
		System.out.println("v1 is value: " + v1);

		// 测试递归
		int v2 = recur(10);
		System.out.println("v2 is value: " + v2);

	}
}
