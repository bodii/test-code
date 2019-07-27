// 编写方法printDouble, 以小数部分p位、整体至少w位来显示浮点数值x。

class PrintDouble {
	public static void printDouble(double x, int p, int w) {
		String s = String.valueOf(x);
		s = s.substring(s.indexOf('.') + 1);  // 小数点部分
		s = s.substring(p - 1, w);
		// x = Double.parseDouble("0." + s);
		System.out.println("0." + s);
	}

	public static void main(String[] args) {
		double d1 = 2.77936;

		System.out.println(d1);
		printDouble(d1, 1, 3);
	}
}
