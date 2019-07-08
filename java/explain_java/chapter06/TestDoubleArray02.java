// 编写一段程序，从头开始依次为元素类型为double型、元素个数为5的数组元素赋值
// 1.1、2.2、3.3、4.4、5.5，并显示

class TestDoubleArray02 {
	public static void main(String[] args) {
		double[] a = new double[5];

		for (int i = 0; i < a.length; i++)
			a[i] = (i + 1) * 1.1;

		for (int i = 0; i < a.length; i++)
			System.out.printf("a[%d] = %.1f\n", i, a[i]);
	}
}
