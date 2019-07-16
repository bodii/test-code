// 连续编号类

class Id {
	static int counter = 0;   // 赋到了哪一个标识编号
	private int id;			  // 标识编号

	// -- 构造函数
	public Id() {
		id = ++counter;
	}

	// -- 获取标识编号
	public int getId() {
		return id;
	}

}
