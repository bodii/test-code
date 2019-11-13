// 将这个类实现为单例(Singletion)的做法
public class SpellChecker2 {
	private final Lexicon dictionary = ...;
	
	private SpellChecker2(...) {}
	public static INSTANCE = new SpellChecker2(...);

	public boolean isValid(String word) { ... }
	public List<String> suggestions(String typo) { ... }
}
