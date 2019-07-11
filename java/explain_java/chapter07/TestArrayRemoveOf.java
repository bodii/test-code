 // 编写方法arrayRmvOf, 返回一个删除了数组a中的元素a[idx]的数组
 
class TestArrayRemoveOf {
	static int[] arrayRmvOf(int[] a, int idx) {
		int[] b = new int[a.length - 1];
		
		for (int i = 0, j = 0; i < a.length; i++) {
			if (i != idx) {
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
		int[] a = {1, 3, 4, 7, 9, 11};

		putArray(a, "a");

		int[] b = arrayRmvOf(a, 2);

		putArray(b, "b");
	}
}
