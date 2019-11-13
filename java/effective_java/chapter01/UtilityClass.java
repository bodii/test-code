// Noninstantiable utility class
// 让这个类包含一个私有构造器，它就不能被实例化
public class UtilityClass {
	// Suppress default constructor for noninstantiability
	private UtilityClass() {
		throw new AssertionError();
	}
}
