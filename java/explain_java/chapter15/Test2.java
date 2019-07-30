// 显示命令行参数和字符串数组

class Test2 {
	static void printStringArray(String[] s) {
		for (int i = 0; i < s.length; i++)
			System.out.println("No." + i + " = " + s[i]);
	}

	public static void main(String[] args) {
		String[] hands = {"石头", "剪刀",  "布"};

		System.out.println("命令行参数");
		printStringArray(args);

		System.out.println("猜拳的手势");
		printStringArray(hands);
	}
}
