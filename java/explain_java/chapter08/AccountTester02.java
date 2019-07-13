// 用于测试银行账户类

class AccountTest02 {
	public static void main(String[] args) {
		// 足立的账户
		Account02 adachi = new Account02("足立幸一", "123456", 1000);
		// 仲田的账户
		Account02 nakata = new Account02("仲田真二", "654321", 200);

		adachi.withdraw(200);
		nakata.deposit(100);

		System.out.println("足立的账户");
		System.out.println("账户名: " + adachi.getName() );
		System.out.println("账户: " + adachi.getNo() );
		System.out.println("可用余额: " + adachi.getBalance() );

		System.out.println("仲田的账户");
		System.out.println("账户名: " + nakata.getName() );
		System.out.println("账户: " + nakata.getNo() );
		System.out.println("可用余额: " + nakata.getBalance() );
	}
}
