// -- 测试动物类
public class AnimalTester {
	public static void main(String[] args) {
		Animal[] a = {
			new Dog("达罗", "柴犬"),
			new Cat("迈克尔", 7),
			new Dog("八公", "秋田犬"),
		};

		for (Animal k : a) {
			k.introduce();			// 调用k引用的实例类型
			System.out.println();	
		}
	}
}
