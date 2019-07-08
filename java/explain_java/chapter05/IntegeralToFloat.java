// 从整数到浮点型的转换(丢失精度的示例)

class IntegeralToFloat {
	public static void main(String[] args) {
		int a = 123456789;
		long b = 1234567890123456789L;

		System.out.println("\ta = " + a);
		System.out.println("(float)a = " + (float)a);


		System.out.println("\tb = " + b);
		System.out.println("(double)b = " + (double)b);
	}
}
