// 银行账户类

class Account02 {
	private String name;   // 账户名
	private String no;     // 账户
	private long balance;  // 可用余额

	// -- 构造函数-- // 
	Account02(String n, String num, long z) {
		name = n; 
		no = num;
		balance = z;
	}

	// -- 确认账户名-- //
	String getName() {
		return name;
	}

	String getNo() {
		return no;
	}

	long getBalance() {
		return balance;
	}

	// -- 存入k日元
	void deposit(long k) {
		balance += k;
	}

	// -- 取出k日元
	void withdraw(long k) {
		balance -= k;
	}
}
