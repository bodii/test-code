// 把类实现为静态工具类的做法
public class SpellChecker {
	private static final Lexicon dictionary = ...;

	private SpellChecker() {} // Noninstantiable

	public static boolean isValid(String word) { ... };
	public static List<String> suggestions(String typo) { ... }

}
