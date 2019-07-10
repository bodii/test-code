// 编写方法minOf,计算数组中元素的最小值

class TestMinOf {
	static int minOf(int[] a) {
		int min = a[0];
		for (int i = 1; i < a.length; i++) 
			if (a[i] < min)
				min = a[i];

		return min;
	}

	public static void main(String[] args) {
		int[] a = { 13, 2, 5, 7, 3, 1, 9 };

		System.out.println("数组a中最小值是: " + minOf(a));
	}
}
