// 标识编号类

package id;

import java.util.Random;

public class RandId {
	private static int counter;
	private int id;

	static {
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
