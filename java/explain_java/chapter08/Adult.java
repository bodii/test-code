// 使用People 类

class Adult {
	public static void main(String[] args) {
		People adult1 = new People("Joke", 183, 69);

		System.out.println("名子: " + adult1.getName());
		System.out.println("身高: " + adult1.getHeight());
		System.out.println("体重: " + adult1.getWeight());
	}
}
