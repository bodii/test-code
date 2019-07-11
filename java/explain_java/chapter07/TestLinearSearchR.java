// 方法linearSearch中，当存在多个与要查找的键值相同的元素时，查找到的是最开头位置的
// 元素，请编写方法linearSearchR,使其查找到的是最末尾位置的元素.

class TestLinearSearchR {
	static int linearSearchR(int[] a, int key) {

		for (int i = a.length - 1; i >= 0; i--)
			if (a[i] == key)
				return i + 1;

		return -1;
	}

	public static void main(String[] args) {
		int[] a = { 1, 3, 5, 7, 6, 9, 4 };

		int key = 3, k = linearSearchR(a, key);
		if (k != -1)
			System.out.println("整数数组a中，" + key + "在倒数第" + k + "位。");
		else
			System.out.println("整数数组a中，" + key + "不存在。");
	}
}
