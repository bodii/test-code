// 银行账户类

public class Account {
	private static int counter = 0;  // 赋到了哪佣标识编号
	private String name;  // 账户名
	private String no;    // 账户
	private long balance; // 可用余额
	private int id;       // 标识编号

	// -- 构造函数
	public Account(String n, String num, long z) {
		name = n;
		no = num;
		balance = z;
		id = ++counter; 
	}

	// -- 确认账户名
	public String getName() {
		return name;
	}

	// -- 确认账户
	public String getNo() {
		return no;
	}

	// -- 确认可用余额
	public long getBalance() {
		return balance;
	}

	// -- 获取标识编号
	public int getId() {
		return id;
	}

	// -- 存入k日元
	public void deposit(long k) {
		balance += k;
	}

	// -- 取出k日元
	public void widthdraw(long k) {
		balance -= k;
	}
}
