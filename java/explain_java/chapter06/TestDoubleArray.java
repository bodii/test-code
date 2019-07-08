// 创建一个元素类型为double型、元素个数为5的数组，并显示其全部元素的值。

class TestDoubleArray {
	public static void main(String[] args) {
		double[] a = new double[5];
		for (int i = 0; i < a.length; i++)
			a[i] = (double) i;

		for (int i = 0; i < a.length; i++)
			System.out.println("a[" + i +"] = " + a[i]);
	}
}
