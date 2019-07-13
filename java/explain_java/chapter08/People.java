// 人的类

class People {
	private String name;
	private int height;
	private int weight;

	People(String name, int height, int weight) {
		this.name = name;
		this.height = height;
		this.weight = weight;
	}

	String getName() {
		return name;
	}

	int getHeight() {
		return height;
	}

	int getWeight() {
		return weight;
	}
}
