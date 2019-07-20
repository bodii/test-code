// -- 会员类
public class Member {
	private String name;
	private int no;
	private int age;

	public Member(String name, int no, int age) {
		this.name = name;
		this.no = no;
		this.age = age;
	}

	public String getName() {
		return name;
	}

	public void print() {
		System.out.printf("No.%d: %s (%d岁)\n", no, name, age);
	}
}
