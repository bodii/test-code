// 编写一段程序，从头开始依次为元素类型为int型、元素个数为5的数组元素赋值5、4、3、2、1，并
// 显示

class TestIntArray01 {
	public static void main(String[] args) {
		int[] a = new int[5];

		for (int i = (a.length - 1), k = 0; i >= 0; i--, k++) {
			a[k] = i + 1;
		}
		

		for (int i = 0; i < a.length; i++)
			System.out.printf("a[%d] = %d\n", i, a[i]);
	}
}
