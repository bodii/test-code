// -- 动物类
public abstract class Animal {
	private String name;
	public Animal(String name) { this.name = name; }
	public abstract void bark();  // 叫
	public abstract String toString(); // 返回字符串表示
	public String getName() { return name; }
	public void introduce() {
		System.out.print(toString());
		bark();
	}
}
