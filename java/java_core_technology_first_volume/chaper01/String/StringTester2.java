public class StringTester2 {
	public static void main(String[] args) {
		String greeting = "Hello";
		int n = greeting.length();
		System.out.println("str length: " + n );

		// 要想得到实际的长度，即码点数量，可以调用：
		int cpCount = greeting.codePointCount(0, greeting.length());
		System.out.println("str code number: " + cpCount);

		// 调用s.charAt(n)将返回位置n的代码单元，n介于0 ~ s.legnth() -1之间:
		char first = greeting.charAt(0);
		System.out.println("str charAt 0 : " + first);

		// 要想得到第i个码点，应该使用下列语句:
		int index = greeting.offsetByCodePoints(0, 3);
		int cp = greeting.codePointAt(index);
		System.out.println("str code print at：" + cp);

		// 生成一个int值的"流", 每个int值对应一个码点。
		int[] codePoints = greeting.codePoints().toArray();

		// 将一个码点数组转换成一个字符串
		String str = new String(codePoints, 0, codePoints.length);

	}
}
