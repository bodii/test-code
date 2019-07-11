// 编写方法arrayClone, 创建并返回一个与数组a相同的数组(元素个数相同，并且所有的元素
// 值都相同的数组)。

class TestArrayClone {
	
	// -- 返回一个与数组a相同的数组(元素个数相同，并且所有元素值都相同的数组)
	static int[] arrayClone(int[] a) {
		return a;
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
		int[] a = { 5, 3, 2, 8, 9, 3, 4, 6 };
		putArray(a, "a");
		
		int[] b  = arrayClone(a);
		putArray(b, "b");

	}
}
