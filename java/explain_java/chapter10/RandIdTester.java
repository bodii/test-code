// 标识编号类

import java.util.Random;

class RandId {
	private static int counter; // 赋到了哪一个标识编号
	private int id;

	static { // 类（静态）初始化器
		Random rand = new Random();
		counter = rand.nextInt(10) * 100;
	}

	// -- 构造函数
	public RandId() {
		id = ++counter;
	}

	// -- 获取标识编号
	public int getId() {
		return id;
	}
}


public class RandIdTester {
	public static void main(String[] args) {
		RandId a = new RandId();
		RandId b = new RandId();
		RandId c = new RandId();

		System.out.println("a的标识编号: " + a.getId());
		System.out.println("b的标识编号: " + b.getId());
		System.out.println("c的标识编号: " + c.getId());
	}
}
