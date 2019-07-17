// 连续编号类

class Id03 {
	static int counter = 0;

	private int id;

	public Id03() {
		id = ++counter;
	}

	public static int getMaxID() {
		return counter;
	}

	public int getId() {
		return id;
	}
}

public class IdTester03 {
	public static void main(String[] args) {
		Id03 a = new Id03();
		Id03 b = new Id03();

		System.out.println("a的标识编号: " + a.getId());
		System.out.println("b的标识编号: " + b.getId());

		System.out.println("Id03.counter = " + Id03.counter);
		System.out.println("a.counter = " + a.counter);
		System.out.println("b.counter = " + b.counter);
		System.out.println("max id = " + Id03.getMaxID());
	}
}
