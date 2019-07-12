// 编写下述的重载方法群，分别计算int型整数x的绝对值、long型整数x的绝对值、
// float型数x的绝对值、double型实数x的绝对值

class TestAbsolute {
	static int absolute(int x) {
		return x < 0 ? -x : x;
	}

	static long absolute(long x) {
		return x < 0 ? -x : x;
	}

	static float absolute(float x) {
		return x < 0 ? -x : x;
	}

	static double absolute(double x) {
		return x < 0 ? -x : x;
	}

	public static void main(String[] args) {

		int a1 = -4;
		long a2 = -4;
		float a3 = -1.4f;
		double a4 = -.4;
		
		System.out.println("int -4的绝对值是:" + absolute(a1));
		System.out.println("long -4的绝对值是:" + absolute(a2));
		System.out.println("float -4的绝对值是:" + absolute(a3));
		System.out.println("double -4的绝对值是:" + absolute(a4));
	}
}
