// -- 狗类
public class Dog extends Animal {
	private String type;	// 狗的品种

	public Dog(String name, String type) {
		super(name); this.type = type;
	}

	public void bark() { System.out.println("汪汪!!"); }

	public String toString() { return type + "的" + getName(); }
}
