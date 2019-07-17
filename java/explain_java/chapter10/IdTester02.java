// 连续编号类

class Id02 {
	static int counter = 0;

	private int id;

	public Id02() {
		id = ++counter;
	}

	public int getId() {
		return id;
	}
}

public class IdTester02 {
	public static void main(String[] args) {
		Id02 a = new Id02();
		Id02 b = new Id02();

		System.out.println("a的标识编号: " + a.getId());
		System.out.println("b的标识编号: " + b.getId());

		System.out.println("Id02.counter = " + Id02.counter);
		System.out.println("a.counter = " + a.counter);
		System.out.println("b.counter = " + b.counter);
	}
}
