// is-A关系和实例的引用（使用方法的参数进行验证)

class TimeAccountTester02 {
	// 谁的可用余额更多呢
	static int compBalance(Account a, Account b) {
		if (a.getBalance() > b.getBalance())
			return 1;
		else if (a.getBalance() < b.getBalance())
			return -1;

		return 0;
	}

	public static void main(String[] args) {
		Account adachi = new Account("足立幸一", "123456", 1000);
		TimeAccount02 nakata = new TimeAccount02("仲田真一", "654321", 200, 500);

		switch(compBalance(adachi, nakata)) {
			case 0: System.out.println("足立和仲田的可用余额相同。"); break;
			case 1: System.out.println("足立的可用余额更多。"); break;
			case -1: System.out.println("仲田的可用余额更多。"); break;
		}
	}
}
