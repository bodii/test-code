// 编写方法aryExchng,交换数组a和b中全部元素的值。

class TestArrayExchang {

	// -- 交换数组a和b中全部元素的值 --//
	static void aryExchng(int[] a, int[] b) {
		int n = a.length < b.length ? a.length : b.length;

		for (int i = 0; i < n; i++) {
			int t = a[i];
			a[i] = b[i];
			b[i] = t;
		}
	}

	// -- 打印数组元素 --//
	static void putArray(int[] a, String arr_name) {
		String s = arr_name + " = { ";
		for (int i = 0; i < a.length; i++) {
			s += a[i] + ", ";
		}
		s += "}";

		System.out.println(s);
	}

	public static void main(String[] args) {
		int[] a = {1, 2, 3, 4, 5, 6, 7};
		int[] b = {5, 4, 3, 2, 1};

		System.out.println("before: ");
		putArray(a, "a");
		putArray(b, "b");

		aryExchng(a, b);

		System.out.println("after: ");
		putArray(a, "a");
		putArray(b, "b");
	}
}
