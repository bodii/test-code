// 编写方法arrayRmvOfN, 返回一个删除了数组a中从元素a[idx]开始的n个元素后的数组

class TestArrayRemoveOfN {
	static int[] arrayRmvOfN(int[] a, int idx, int n) {
		int[] b = new int[n];

		for (int i = 0, j = 0; i < a.length; i++) {
			if (i < idx || i >= (idx + n)) {
				b[j] = a[i];
				j++;
			}
			
		}

		return b;
	}

	// -- 打印数组元素
	static void putArray(int[] a, String arr_name) {
		String s = arr_name + " =  { ";
		for (int i = 0; i < a.length; i++) {
			s += a[i] + ", ";
		}
		s += "}";

		System.out.println(s);
	}

	public static void main(String[] args) {
		int[] a = { 1, 3, 4, 7, 9, 11 };
		putArray(a , "a");

		int[] b = arrayRmvOfN(a, 1, 3);
		putArray(b, "b");

	}
}
