// 操作两个人的银行账户数据的程序

class Accounts {
	public static void main(String[] args) {
		String adachiAccountName		= "足立幸一"; 
		String adachiAccountNo			= "123456"; 
		long adachiAccountBalance		= 1000;

		String nakataAccountName		= "仲田真二";
		String nakataAccountNo			= "654321";
		long   nakataAccountBalance		= 200;

		adachiAccountBalance -= 200;
		nakataAccountBalance += 100;

		System.out.println("足立的账户");
		System.out.println("账户名: " + adachiAccountName);
		System.out.println("账户: " + adachiAccountNo);
		System.out.println("可用余额: " + adachiAccountBalance);

		System.out.println("仲田的账户");
		System.out.println("账户名: " + nakataAccountName);
		System.out.println("账户: " + nakataAccountNo);
		System.out.println("可用余额: " + nakataAccountBalance);
		
	}
}
