// -- 会员类
public class SpecialMember extends Member {
	private String privilege;  // 优惠

	public SpecialMember(String name, int no, int age, String privilege) {
		super(name, no, age); 
		this.privilege = privilege;
	}

	@Override public void print() {
		super.print();
		System.out.println("优惠: " + privilege);
	}
}
