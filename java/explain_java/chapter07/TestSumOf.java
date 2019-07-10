// 编写方法sumOf,计算数组a中全部元素的总和

class TestSumOf {
	static int sumOf(int[] arr) {
		int sum = 0;
		for (int i = 0; i < arr.length; i++) 
			sum += arr[i];

		return sum;
	}

	public static void main(String[] args) {
		int[] a = {1, 2, 3, 4, 5, 6};
		System.out.println("数组a中全部元素的总和是: " + sumOf(a));
	}
}
