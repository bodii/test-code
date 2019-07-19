// 宠物类

class Pet {
	private String name;		// 宠物名字
	private String masterName;	// 主人的名字

	// 构造函数
	public Pet(String name, String masterName) {
		this.name = name;
		this.masterName = masterName;
	}

	// 确认宠物的名字
	public String getName() { return name; }

	// 确认主人的名字
	public String getMasterName() { return masterName; }

	// 自我介绍
	public void introduce() {
		System.out.println("我的名字是" + name + "! " );
		System.out.println("主人的名字是" + masterName + "! " );
	}
}

class RobotPet extends Pet {

	// 构造函数
	public RobotPet(String name, String masterName) {
		super(name, masterName);
	}

	// 自我介绍
	public void introduce() {
		System.out.println("我是机器人。名字是" + getName() + " 。");
		System.out.println("我是主人。名字是" + getMasterName() + " 。");
	}

	// 做家务
	public void work(int sw) {
		switch(sw) {
			case 0: System.out.println("打扫。"); break;
			case 1: System.out.println("洗衣服。"); break;
			case 2: System.out.println("做饭。"); break;
		}
	}
}
