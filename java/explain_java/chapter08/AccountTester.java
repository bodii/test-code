// 用于测试银行账户类的类

class AccountTester {
	public static void main(String[] args) {
		Account adachi = new Account();  // 足立的账户
		Account nakata = new Account();	 // 仲田的账户

		adachi.name			= "足立幸一";
		adachi.no			= "123456";
		adachi.balance		= 1000;

		nakata.name			= "仲田真二";
		nakata.no			= "654321";
		nakata.balance		= 200;

		adachi.balance -= 200;
		nakata.balance -= 100;

		System.out.println("足立的账户");
		System.out.println("账户名: " + adachi.name);
		System.out.println("账户: " + adachi.no);
		System.out.println("可用余额: " + adachi.balance);
		System.out.println();
		nakata.balance -= 100;

		System.out.println("足立的账户");
		System.out.println("账户名: " + nakata.name);
		System.out.println("账户: " + nakata.no);
		System.out.println("可用余额: " + nakata.balance);
		System.out.println();
	}
}
