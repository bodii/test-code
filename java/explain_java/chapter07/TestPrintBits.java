// 方法printBits用于显示int型整数值的各个位的内容。

class TestPrintBits {
	static void printBits(byte x) {
		System.out.println("byte x: " + x);
	}

	static void printBits(short x) {
		System.out.println("short x: " + x);
	}
  
	static void printBits(int x) {
		System.out.println("int x: " + x);
	}

	static void printBits(long x) {
		System.out.println("long x: " + x);
	}

	public static void main(String[] args) {
		byte x1 = (byte)(1245);
		printBits(x1);

		short x2 = 21;
		printBits(x2);

		int x3 = 21;
		printBits(x3);

		long x4 = 21;
		printBits(x4);
	}
}
