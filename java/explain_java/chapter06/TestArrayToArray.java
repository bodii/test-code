// 将数组a中的全部元素倒序复制到数组b中。

class TestArrayToArray {
	public static void main(String[] args) {
		int[] a = { 42, 35, 85, 2, -7 };
		int[] b = { 0, 0, 0, 0, 0 };

		for (int i = 0; i < a.length; i++) {
			for (int j =  i + 1; j < a.length; j++) {
				if (a[i] < a[j]) {
					int t  = a[i];
					a[i] = a[j];
					a[j] = t;
				}
			}
			b[i] = a[i];
		}

		for (int i = 0; i < a.length; i++) {
			System.out.printf("a[%2$d] = %1$d, ", a[i], i);
			System.out.printf("b[%2$d] = %1$d\n", b[i], i);
		}
	}
}
