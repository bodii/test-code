// 比较拼写相同的字符串的引用目标
class StringLiteral {
	public static void main(String[] args) {
		String s1 = "ABC";
		String s2 = "ABC";

		if (s1 == s2 )
			System.out.println("s1和s2引用相同的字符串。");
		else
			System.out.println("s1和s2引用不同的字符串。");
	}
}
