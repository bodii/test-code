// is－A关系和实例的引用

class TimeAccountTester01 {
	public static void main(String[] args) {
		Account adachi = new Account("足立幸一", "123456", 1000);
		TimeAccount02 nakata = new TimeAccount02("中田真二", "654321", 200, 500);

		Account x;		// 类类型变量
		x = adachi;		// 可以引用自身类型的实例
		x = nakata;		// 也可以引用下位类类型的实例

		System.out.println("x的可用余额: " + x.getBalance());

		TimeAccount02 y;
		y = nakata;

		System.out.println("y的可用余额: " + y.getBalance());
		System.out.println("y的定期余额: " + y.getTimeBalance());
	}
}
