// Id类测试

public class IdTester {
	public static void main(String[] args) {
		Id a = new Id();
		Id b = new Id();

		System.out.println("a的标识编号: " + a.getId());
		System.out.println("b的标识编号: " + b.getId());
		
		System.out.println("Id.counter =  " + Id.counter);
		System.out.println("a.counter =  " + a.counter);
		System.out.println("b.counter =  " + b.counter);

	}
}
