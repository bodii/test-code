class TestMinMaxTester {

	public static void main(String[] args) {
		int a = 10, b = 5, c = 3;
		int[] arr = {10, 15, 11 };

		System.out.println("两个数中的最小数：" + TestMinMax.min(a, b));
		System.out.println("两个数中的最大数：" + TestMinMax.max(a, b));
		
		System.out.println("三个数中的最小数：" + TestMinMax.min(a, b, c));
		System.out.println("三个数中的最大数：" + TestMinMax.max(a, b, c));

		System.out.println("数组中的最小数：" + TestMinMax.min(arr));
		System.out.println("数组中的最大数：" + TestMinMax.max(arr));
	}
}
