// 编写方法arraySrchIdx, 返回一个数组，该数组中从头开始依次保存数组a中值为x的全部元素的索引

import java.util.ArrayList;

class TestArraySearchIndex {

	// -- 返回数组为x的全部索引
	static int[] arraySrchIdx(int[] a, int x) {

		ArrayList r = new ArrayList();
		for (int i = 0; i < a.length; i++) {
			if (a[i] == x) {
				r.add(i);
			}
		}

		int[] v = r.toArray(new int[r.size()]);
		
		return v;
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
		int[] a = { 1, 5, 4, 8, 5, 5, 7};
		putArray(a, "a");

		int[] b = arraySrchIdx(a, 5);
		putArray(b, "b");
	}
}
